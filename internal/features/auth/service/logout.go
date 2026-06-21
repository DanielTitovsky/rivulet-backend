package auth_service

import (
	"context"
	"fmt"
)

func (s *AuthServise) Logout(
	ctx context.Context,
	refreshTokenString string,
) error {
	err := s.startLogoutUserTx(ctx, refreshTokenString)

	if err != nil {
		return err
	}

	return nil
}

func (s *AuthServise) startLogoutUserTx(
	ctx context.Context,
	refreshTokenString string,
) error {
	claims, err := s.TokenService.ValidateToken(refreshTokenString, "refresh")
	if err != nil {
		return fmt.Errorf("validate refresh token: %w", err)
	}

	savedRefreshToken, err := s.TokenService.GetRefreshToken(ctx, refreshTokenString)
	if err != nil {
		return fmt.Errorf("get refresh token: %w", err)
	}

	if savedRefreshToken.Id != claims.Id {
		return fmt.Errorf("invalid refresh token")
	}

	if savedRefreshToken.UserId != claims.UserId {
		return fmt.Errorf("invalid refresh token user")
	}

	err = s.TokenService.RemoveToken(ctx, savedRefreshToken.Id)
	if err != nil {
		return fmt.Errorf("remove refresh token: %w", err)
	}

	return nil
}
