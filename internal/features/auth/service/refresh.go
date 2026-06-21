package auth_service

import (
	"context"
	"fmt"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
)

func (s *AuthServise) Refresh(
	ctx context.Context,
	refreshTokenString string,
) (domain.User, *domain.Token, *domain.Token, error) {
	user, accessToken, refreshToken, err := s.startRefreshUserTx(ctx, refreshTokenString)

	if err != nil {
		return domain.User{}, nil, nil, err
	}

	return user, accessToken, refreshToken, nil
}

func (s *AuthServise) startRefreshUserTx(
	ctx context.Context,
	refreshTokenString string,
) (domain.User, *domain.Token, *domain.Token, error) {
	claims, err := s.TokenService.ValidateToken(refreshTokenString, "refresh")
	if err != nil {
		return domain.User{}, nil, nil, fmt.Errorf("validate refresh token: %w", err)
	}

	savedRefreshToken, err := s.TokenService.GetRefreshToken(ctx, refreshTokenString)
	if err != nil {
		return domain.User{}, nil, nil, fmt.Errorf("get refresh token: %w", err)
	}

	if savedRefreshToken.Id != claims.Id {
		return domain.User{}, nil, nil, fmt.Errorf("invalid refresh token")
	}

	if savedRefreshToken.UserId != claims.UserId {
		return domain.User{}, nil, nil, fmt.Errorf("invalid refresh token user")
	}

	user, err := s.UserService.GetUser(ctx, claims.UserId)
	if err != nil {
		return domain.User{}, nil, nil, fmt.Errorf("get user: %w", err)
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
