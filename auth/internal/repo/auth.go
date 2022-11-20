package repo

import (
	"context"
	"encoding/json"
	"fmt"
	conf "github.com/Skijetler/alphinium/auth/internal/config"
	"github.com/Skijetler/alphinium/auth/internal/usecase"
	"github.com/Skijetler/alphinium/pkg/ent"
	"github.com/Skijetler/alphinium/pkg/ent/user"
	"github.com/frankenbeanies/randhex"
	"github.com/go-redis/redis/v9"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"math/rand"
	"time"
)

type authRepo struct {
	repo *Repo
	log  *logrus.Logger
}

// NewAuthRepo .
func NewAuthRepo(repo *Repo, logger *logrus.Logger) usecase.AuthRepo {
	rand.Seed(time.Now().UnixNano())

	return &authRepo{
		repo: repo,
		log:  logger,
	}
}

func (r *authRepo) WithTx(ctx context.Context, fn func(tx *ent.Tx) error) error {
	tx, err := r.repo.db.Tx(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if v := recover(); v != nil {
			_ = tx.Rollback()
			panic(v)
		}
	}()

	if err = fn(tx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			r.log.WithContext(ctx).Errorf("Rolling back transaction error: %v", rerr)
		}
		return err
	}
	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (r *authRepo) WithTxRedis(ctx context.Context, fn func(pipe redis.Pipeliner) error) error {
	err := r.repo.sessionDB.Watch(ctx, func(tx *redis.Tx) error {
		_, err := tx.TxPipelined(ctx, fn)
		return err
	})
	if err != nil {
		return err
	}

	return nil
}

func (r *authRepo) SaveUser(ctx context.Context, u *usecase.User) (*usecase.User, error) {
	var userModel *ent.User
	var userMetadata *ent.UserMetadata

	if err := r.WithTx(ctx, func(tx *ent.Tx) error {
		var err error

		if userModel, err = tx.User.
			Create().
			SetName(u.Name).
			SetEmail(u.Email).
			SetPassword(u.Password).
			SetRegistrationDate(time.Now().UTC()).
			Save(ctx); err != nil {
			return err
		}

		if userMetadata, err = tx.UserMetadata.
			Create().
			SetColor(randhex.New().String()).
			SetTitle(u.Title).
			SetGender(u.Gender).
			SetLastOnline(userModel.RegistrationDate).
			SetUser(userModel).
			Save(ctx); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return &usecase.User{
		ID:       userModel.ID,
		Name:     userModel.Name,
		Title:    userMetadata.Title,
		Gender:   userMetadata.Gender,
		Email:    userModel.Email,
		Password: userModel.Password,
		Disabled: userModel.Disabled,
	}, nil
}

func (r *authRepo) GetUserByUsername(ctx context.Context, username string) (*usecase.User, error) {
	u, err := r.repo.db.User.Query().Where(user.Name(username)).Only(ctx)
	if err != nil {
		return nil, err
	}

	m, err := u.QueryMetadata().Only(ctx)
	if err != nil {
		return nil, err
	}

	return &usecase.User{
		ID:       u.ID,
		Name:     u.Name,
		Title:    m.Title,
		Gender:   m.Gender,
		Email:    u.Email,
		Password: u.Password,
		Disabled: u.Disabled,
	}, nil
}

func (r *authRepo) GetUserByEmail(ctx context.Context, email string) (*usecase.User, error) {
	u, err := r.repo.db.User.Query().Where(user.Email(email)).Only(ctx)
	if err != nil {
		return nil, err
	}

	m, err := u.QueryMetadata().Only(ctx)
	if err != nil {
		return nil, err
	}

	return &usecase.User{
		ID:       u.ID,
		Name:     u.Name,
		Title:    m.Title,
		Gender:   m.Gender,
		Email:    u.Email,
		Password: u.Password,
		Disabled: u.Disabled,
	}, nil
}

func (r *authRepo) CheckUserIsDisabled(ctx context.Context, userId uint64) (bool, error) {
	u, err := r.repo.db.User.Get(ctx, userId)
	if err != nil {
		return false, err
	}

	return u.Disabled, nil
}

func (r *authRepo) SaveSession(ctx context.Context, s *usecase.Session) (string, error) {
	var sessionId string
	var isKeyExists bool
	c := conf.GetConfig("")

	for i := 0; i < c.UUID.Iterations; i++ {
		u, err := uuid.NewUUID()
		if err != nil {
			return "", err
		}

		sessionId = u.String()

		isKeyExists, err := r.repo.sessionDB.Do(ctx, "EXISTS", sessionId).Result()
		if err != nil {
			return "", err
		}

		if isKeyExists == false {
			sessionSerialized, err := json.Marshal(s)
			if err != nil {
				return "", err
			}

			if err := r.WithTxRedis(ctx, func(pipe redis.Pipeliner) error {
				pipe.Set(ctx, sessionId, string(sessionSerialized), c.TokenMaker.RefreshTtl)
				return nil
			}); err != nil {
				return "", err
			}
			break
		}
	}
	if isKeyExists {
		return "", fmt.Errorf("can not generate session id")
	}

	return sessionId, nil
}

func (r *authRepo) UpdateSessionTTLById(ctx context.Context, sessionId string) (bool, error) {
	var keysCount int64

	if err := r.WithTxRedis(ctx, func(pipe redis.Pipeliner) error {
		var err error
		keysCount, err = pipe.Touch(ctx, sessionId).Result()
		return err
	}); err != nil {
		return false, err
	}

	if keysCount == 0 {
		return false, nil
	}

	return true, nil
}

func (r *authRepo) GetSessionById(ctx context.Context, sessionId string) (*usecase.Session, error) {
	var sessionString string
	var s *usecase.Session

	if err := r.WithTxRedis(ctx, func(pipe redis.Pipeliner) error {
		var err error
		sessionString, err = pipe.Get(ctx, sessionId).Result()
		return err
	}); err != nil {
		return nil, err
	}

	err := json.Unmarshal([]byte(sessionString), &s)
	if err != nil {
		return nil, err
	}

	return s, nil
}
