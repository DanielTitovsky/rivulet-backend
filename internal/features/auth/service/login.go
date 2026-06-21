package auth_service

import (
	"context"
	"fmt"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
)

func (s *AuthServise) Login(
	ctx context.Context,
	email string,
	password string,
) (domain.User, *domain.Token, *domain.Token, error) {
	user, accessToken, refreshToken, err := s.startLoginUserTx(ctx, email, password)

	if err != nil {
		return domain.User{}, nil, nil, err
	}

	return user, accessToken, refreshToken, nil
}

func (s *AuthServise) startLoginUserTx(
	ctx context.Context,
	email string,
	password string,
) (domain.User, *domain.Token, *domain.Token, error) {
	user, err := s.UserService.GetUserByEmail(ctx, email)
	if err != nil {
		return domain.User{}, nil, nil, fmt.Errorf("invalid email or password")
	}

	if user.Password == nil {
		return domain.User{}, nil, nil, fmt.Errorf("invalid email or password")
	}

	if !user.PasswordComparison(password, *user.Password) {
		return domain.User{}, nil, nil, fmt.Errorf("invalid email or password")
	}

	accessToken, refreshToken, err := s.TokenService.GenerateTokens(
		AccessExpired,
		RefreshExpired,
		user,
	)
	if err != nil {
		return domain.User{}, nil, nil, fmt.Errorf("generate tokens: %w", err)
	}

	_, err = s.TokenService.SaveRefreshToken(ctx, *refreshToken)
	if err != nil {
		return domain.User{}, nil, nil, fmt.Errorf("save refresh token: %w", err)
	}

	return user, accessToken, refreshToken, nil
}
