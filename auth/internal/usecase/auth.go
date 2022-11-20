package usecase

import (
	"context"
	v1 "github.com/Skijetler/alphinium/auth/api/v1"
	"github.com/Skijetler/alphinium/auth/internal/pkg/hash"
	"github.com/Skijetler/alphinium/auth/internal/pkg/paseto"
	"github.com/sirupsen/logrus"
)

// User is a user model.
type User struct {
	ID       uint64
	Name     string
	Title    string
	Gender   string
	Email    string
	Password string
	Disabled bool
}

// Session is a session model
type Session struct {
	UserId uint64 `json:"userId" binding:"required"`
}

// AuthRepo is an Auth repo.
type AuthRepo interface {
	SaveUser(context.Context, *User) (*User, error)
	GetUserByUsername(context.Context, string) (*User, error)
	GetUserByEmail(context.Context, string) (*User, error)
	CheckUserIsDisabled(context.Context, uint64) (bool, error)
	SaveSession(context.Context, *Session) (string, error)
	UpdateSessionTTLById(context.Context, string) (bool, error)
	GetSessionById(context.Context, string) (*Session, error)
}

// AuthUsecase is an Auth usecase.
type AuthUsecase struct {
	repo AuthRepo
	log  *logrus.Logger

	hasher     hash.PasswordHasher
	tokenMaker paseto.TokenMaker
}

// NewAuthUsecase new an Auth usecase.
func NewAuthUsecase(repo AuthRepo, logger *logrus.Logger, hasher hash.PasswordHasher, tokenMaker paseto.TokenMaker) *AuthUsecase {
	return &AuthUsecase{repo: repo, log: logger, hasher: hasher, tokenMaker: tokenMaker}
}

// NewTokens generate new access & refresh paseto tokens
func (uc *AuthUsecase) NewTokens(ctx context.Context, sessionId string) (*v1.Tokens, error) {
	uc.log.WithContext(ctx).Infof("NewTokens: %v", sessionId)

	access, err := uc.tokenMaker.NewAccessToken(sessionId)
	if err != nil {
		return nil, internalErr(err)
	}

	refresh, err := uc.tokenMaker.NewRefreshToken(sessionId)
	if err != nil {
		return nil, internalErr(err)
	}

	return &v1.Tokens{
		AccessToken:  access,
		RefreshToken: refresh,
	}, nil
}

// CreateUser save new user to db
func (uc *AuthUsecase) CreateUser(ctx context.Context, u *User) (*User, error) {
	uc.log.WithContext(ctx).Infof("CreateUser: %v", u.Name)
	u.Password = uc.hasher.Hash(u.Password)
	saved, err := uc.repo.SaveUser(ctx, u)
	if err != nil {
		uc.log.WithContext(ctx).Errorf("CreateUser error: %v", err)
		return nil, ErrUserAlreadyExists
	}
	return saved, nil
}

// ComparePassword compare pass with the one saved in db
func (uc *AuthUsecase) ComparePassword(ctx context.Context, username, pass string) (*User, error) {
	uc.log.WithContext(ctx).Infof("ComparePassword: %v", username)

	model, err := uc.repo.GetUserByUsername(ctx, username)
	if err != nil {
		uc.log.WithContext(ctx).Errorf("ComparePassword error: %v", err)
		return nil, ErrUserNotFound
	}

	if ok := uc.hasher.Compare(model.Password, pass); !ok {
		return nil, ErrWrongPassword
	}

	if model.Disabled {
		return nil, ErrUserIsDisabled
	}

	return model, nil
}

// CreateSession save new session to db, returns session id
func (uc *AuthUsecase) CreateSession(ctx context.Context, s *Session) (string, error) {
	uc.log.WithContext(ctx).Infof("CreateSession: %v", s.UserId)
	sessionId, err := uc.repo.SaveSession(ctx, s)
	if err != nil {
		uc.log.WithContext(ctx).Errorf("CreateSession error: %v", err)
		return "", internalErr(err)
	}
	return sessionId, nil
}

// GetIdFromRefresh parse refresh token & return session id
func (uc *AuthUsecase) GetIdFromRefresh(ctx context.Context, refresh string) (string, error) {
	uc.log.WithContext(ctx).Infof("GetIdFromRefresh: %v", refresh)

	id, err := uc.tokenMaker.ParseRefreshToken(refresh)
	if err != nil {
		uc.log.WithContext(ctx).Errorf("GetIdFromRefresh error: %v", err)
		return "", ErrInvalidToken
	}

	session, err := uc.repo.GetSessionById(ctx, id)
	if err != nil {
		uc.log.WithContext(ctx).Errorf("GetIdFromRefresh error: %v", err)
		return "", ErrInvalidSession
	}

	userIsDisabled, err := uc.repo.CheckUserIsDisabled(ctx, session.UserId)
	if err != nil {
		uc.log.WithContext(ctx).Errorf("GetIdFromRefresh error: %v", err)
		return "", ErrUserNotFound
	}

	if userIsDisabled {
		uc.log.WithContext(ctx).Errorf("GetIdFromRefresh error: %v", err)
		return "", ErrUserIsDisabled
	}

	return id, nil
}

// Identify identifies user
func (uc *AuthUsecase) Identify(ctx context.Context, access string) (*Session, error) {
	sessionId, err := uc.tokenMaker.ParseAccessToken(access)
	if err != nil {
		uc.log.WithContext(ctx).Errorf("Invalid token: %v", access)
		return nil, ErrInvalidToken
	}

	uc.log.WithContext(ctx).Infof("Identity: %v", sessionId)

	s, err := uc.repo.GetSessionById(ctx, sessionId)
	if err != nil {
		return nil, internalErr(err)
	}
	return s, nil
}

func (uc *AuthUsecase) UpdateUserSessionTTL(ctx context.Context, sessionId string) (bool, error) {
	updated, err := uc.repo.UpdateSessionTTLById(ctx, sessionId)
	if err != nil {
		uc.log.WithContext(ctx).Errorf("UpdateUserSessionTTL error: %v", err)
		return false, internalErr(err)
	}

	return updated, err
}
