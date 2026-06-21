package auth_service

import (
	"context"
	"fmt"
	"time"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
)

const (
	AccessExpired  = 15 * time.Minute
	RefreshExpired = 30 * 24 * time.Hour
)

type RegisterResult struct {
	User         domain.User
	AccessToken  *domain.Token
	RefreshToken *domain.Token
}

func (s *AuthServise) Register(
	ctx context.Context,
	user domain.User,
) (domain.User, *domain.Token, *domain.Token, error) {
	var createdUser domain.User
	var accessToken *domain.Token
	var refreshToken *domain.Token

	createdUser, accessToken, refreshToken, err := s.startRegisterUserTx(ctx, user)

	if err != nil {
		return domain.User{}, nil, nil, err
	}

	return createdUser, accessToken, refreshToken, nil
}

func (s *AuthServise) startRegisterUserTx(ctx context.Context, user domain.User) (domain.User, *domain.Token, *domain.Token, error) {
	createdUser, err := s.UserService.CreateUser(ctx, user)
	if err != nil {
		return domain.User{}, nil, nil, fmt.Errorf("create user: %w", err)
	}

	accessToken, refreshToken, err := s.TokenService.GenerateTokens(
		AccessExpired,
		RefreshExpired,
		createdUser,
	)
	if err != nil {
		return domain.User{}, nil, nil, fmt.Errorf("generate tokens: %w", err)
	}

	_, err = s.TokenService.SaveRefreshToken(ctx, *refreshToken)
	if err != nil {
		return domain.User{}, nil, nil, fmt.Errorf("save refresh token: %w", err)
	}

	return createdUser, accessToken, refreshToken, nil
}
