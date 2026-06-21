package auth_service

import (
	"context"
	"fmt"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
)

func (s *AuthServise) Me(
	ctx context.Context,
	accessTokenString string,
) (domain.User, error) {
	user, err := s.getCurrentUser(ctx, accessTokenString)

	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (s *AuthServise) getCurrentUser(
	ctx context.Context,
	accessTokenString string,
) (domain.User, error) {
	claims, err := s.TokenService.ValidateToken(accessTokenString, "access")
	if err != nil {
		return domain.User{}, fmt.Errorf("validate access token: %w", err)
	}

	user, err := s.UserService.GetUser(ctx, claims.UserId)
	if err != nil {
		return domain.User{}, fmt.Errorf("get user: %w", err)
	}

	return user, nil
}
