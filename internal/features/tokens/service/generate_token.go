package tokens_service

import (
	"fmt"
	"time"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

func (s *TokensServise) GenerateTokens(
	accessExpires time.Duration,
	refreshExpires time.Duration,
	user domain.User,
) (*domain.Token, *domain.Token, error) {
	now := time.Now()

	accessToken := domain.NewTokenUninitialized(
		uuid.New(),
		user.Id,
		now.Add(accessExpires),
		now,
	)

	refreshToken := domain.NewTokenUninitialized(
		uuid.New(),
		user.Id,
		now.Add(refreshExpires),
		now,
	)

	userEmail := ""

	if user.Email != nil {
		userEmail = *user.Email
	}

	accessClaims := domain.TokenClaims{
		Id:        accessToken.Id,
		UserId:    user.Id,
		UserEmail: userEmail,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: accessToken.ExpiresAt.Unix(),
			IssuedAt:  now.Unix(),
		},
	}

	refreshClaims := domain.TokenClaims{
		Id:        refreshToken.Id,
		UserId:    user.Id,
		UserEmail: userEmail,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: refreshToken.ExpiresAt.Unix(),
			IssuedAt:  now.Unix(),
		},
	}

	accessJwt := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	refreshJwt := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)

	accessTokenString, err := accessJwt.SignedString(s.accessSecret)
	if err != nil {
		return nil, nil, fmt.Errorf("sign access token: %w", err)
	}

	refreshTokenString, err := refreshJwt.SignedString(s.refreshSecret)
	if err != nil {
		return nil, nil, fmt.Errorf("sign refresh token: %w", err)
	}

	accessToken.TokenString = accessTokenString
	refreshToken.TokenString = refreshTokenString

	return accessToken, refreshToken, nil
}
