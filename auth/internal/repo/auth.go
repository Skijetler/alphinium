package repo

import (
	"context"
	"github.com/Skijetler/alphinium/auth/internal/usecase"
	"github.com/Skijetler/alphinium/pkg/ent"
	"github.com/frankenbeanies/randhex"
	"github.com/go-redis/redis/v9"
	"github.com/sirupsen/logrus"
	"math/rand"
	"time"
)

type userRepo struct {
	repo *Repo
	log  *logrus.Logger
}

// NewAuthRepo .
func NewAuthRepo(repo *Repo, logger *logrus.Logger) usecase.UserRepo {
	rand.Seed(time.Now().UnixNano())

	return &userRepo{
		repo: repo,
		log:  logger,
	}
}

func (r *userRepo) WithTx(ctx context.Context, fn func(tx *ent.Tx) error) error {
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

func (r *userRepo) WithTxRedis(ctx context.Context, fn func(tx *redis.Tx) error) error {
	err := r.repo.sessionDB.Watch(ctx, fn)
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepo) Save(ctx context.Context, u *usecase.User) (*usecase.User, error) {
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
