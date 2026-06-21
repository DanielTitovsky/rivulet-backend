package auth_service

import (
	"context"
	"errors"
	"fmt"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
	"github.com/jackc/pgx/v5"
)

func (s *AuthServise) OAuthLogin(
	ctx context.Context,
	oauthUser domain.OAuthUser,
) (domain.User, *domain.Token, *domain.Token, error) {
	user, accessToken, refreshToken, err := s.startOAuthLoginTx(ctx, oauthUser)

	if err != nil {
		return domain.User{}, nil, nil, err
	}

	return user, accessToken, refreshToken, nil
}

func (s *AuthServise) startOAuthLoginTx(
	ctx context.Context,
	oauthUser domain.OAuthUser,
) (domain.User, *domain.Token, *domain.Token, error) {
	if !oauthUser.EmailVerified {
		return domain.User{}, nil, nil, fmt.Errorf("oauth email is not verified")
	}

	if oauthUser.ProviderUserId == "" {
		return domain.User{}, nil, nil, fmt.Errorf("oauth provider user id is empty")
	}

	if oauthUser.ProviderUserEmail == "" {
		return domain.User{}, nil, nil, fmt.Errorf("oauth email is empty")
	}

	userId, err := s.AuthRepository.GetOAuthAccountUserId(
		ctx,
		oauthUser.Provider,
		oauthUser.ProviderUserId,
	)

	if err == nil {
		user, err := s.UserService.GetUser(ctx, userId)
		if err != nil {
			return domain.User{}, nil, nil, fmt.Errorf("get oauth user: %w", err)
		}

		return s.createOAuthTokens(ctx, user)
	}

	if !errors.Is(err, pgx.ErrNoRows) {
		return domain.User{}, nil, nil, fmt.Errorf("get oauth account user id: %w", err)
	}

	name := oauthUser.Name
	if name == "" {
		name = oauthUser.ProviderUserEmail
	}

	user, err := s.UserService.GetOrCreateOAuthUser(
		ctx,
		oauthUser.ProviderUserEmail,
		name,
	)
	if err != nil {
		return domain.User{}, nil, nil, fmt.Errorf("get or create oauth user: %w", err)
	}

	err = s.AuthRepository.CreateOAuthAccount(ctx, user.Id, oauthUser)
	if err != nil {
		return domain.User{}, nil, nil, fmt.Errorf("create oauth account: %w", err)
	}

	return s.createOAuthTokens(ctx, user)
}

func (s *AuthServise) createOAuthTokens(
	ctx context.Context,
	user domain.User,
) (domain.User, *domain.Token, *domain.Token, error) {
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
