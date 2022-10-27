package service

import (
	v1 "github.com/Skijetler/alphinium/auth/api/v1"
	"github.com/Skijetler/alphinium/auth/internal/usecase"
	"golang.org/x/net/context"
)

// AuthService is an auth service.
type AuthService struct {
	v1.UnimplementedAuthServer

	uc *usecase.AuthUsecase
}

// NewAuthService new an auth service.
func NewAuthService(uc *usecase.AuthUsecase) *AuthService {
	return &AuthService{uc: uc}
}

// SignUp implements auth.SignUp.
func (s *AuthService) SignUp(ctx context.Context, in *v1.SignUpRequest) (*v1.SignUpReply, error) {
	model, err := s.uc.CreateUser(ctx, &usecase.User{
		Name:     in.Username,
		Title:    in.Title,
		Gender:   in.Gender,
		Email:    in.Email,
		Password: in.Password,
	})
	if err != nil {
		return nil, err
	}

	sessionId, err := s.uc.CreateSession(ctx, &usecase.Session{
		UserId: model.ID,
	})
	if err != nil {
		return nil, err
	}

	tokens, err := s.uc.NewTokens(ctx, sessionId)
	if err != nil {
		return nil, err
	}

	return &v1.SignUpReply{Tokens: tokens}, nil
}

// SignIn implements auth.SignIn.
func (s *AuthService) SignIn(ctx context.Context, in *v1.SignInRequest) (*v1.SignInReply, error) {
	model, err := s.uc.ComparePassword(ctx, in.Username, in.Password)
	if err != nil {
		return nil, err
	}

	sessionId, err := s.uc.CreateSession(ctx, &usecase.Session{
		UserId: model.ID,
	})
	if err != nil {
		return nil, err
	}

	tokens, err := s.uc.NewTokens(ctx, sessionId)
	if err != nil {
		return nil, err
	}

	return &v1.SignInReply{Tokens: tokens}, nil
}

// RefreshToken implements auth.RefreshToken.
func (s *AuthService) RefreshToken(ctx context.Context, in *v1.RefreshTokenRequest) (*v1.RefreshTokenReply, error) {
	sessionId, err := s.uc.GetIdFromRefresh(ctx, in.RefreshToken)
	if err != nil {
		return nil, err
	}

	updated, err := s.uc.UpdateUserSessionTTL(ctx, sessionId)
	if err != nil || updated == false {
		return nil, err
	}

	tokens, err := s.uc.NewTokens(ctx, sessionId)
	if err != nil {
		return nil, err
	}

	return &v1.RefreshTokenReply{Tokens: tokens}, nil
}

// Identify implements auth.Identify.
func (s *AuthService) Identify(ctx context.Context, in *v1.IdentifyRequest) (*v1.IdentifyReply, error) {
	session, err := s.uc.Identify(ctx, in.AccessToken)
	if err != nil {
		return nil, err
	}

	return &v1.IdentifyReply{UserId: session.UserId}, nil
}
