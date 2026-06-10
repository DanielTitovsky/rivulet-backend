package token_service

import (
	"fmt"
	"time"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
	"github.com/golang-jwt/jwt"
	uuid "github.com/google/uuid"
)

func (th *TokenService) GenerateTokens(accessExpires time.Duration, refreshExpires time.Duration, user domain.User) (*domain.Token, *domain.Token, error) {

	accessToken := domain.NewTokenUninitialized(
		uuid.New(),
		user.Id,
		time.Now().Add(accessExpires),
		time.Now(),
	)

	accessClaims := domain.TokenClaims{
		Id:        accessToken.Id,
		UserId:    user.Id,
		UserEmail: *user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(accessExpires).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	refreshToken := domain.NewTokenUninitialized(
		uuid.New(),
		user.Id,
		time.Now().Add(accessExpires),
		time.Now(),
	)

	refreshClaims := domain.TokenClaims{
		Id:        refreshToken.Id,
		UserId:    user.Id,
		UserEmail: *user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(refreshExpires).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	accessTokenString := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	refreshTokenString := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)

	accessTokenSigned, err := accessTokenString.SignedString(th.tokenAccessSecret)

	if err != nil {
		return nil, nil, fmt.Errorf("Invalid access signed: %w", err)
	}

	refreshTokenSigned, err := refreshTokenString.SignedString(th.tokenRefreshSecret)

	if err != nil {
		return nil, nil, fmt.Errorf("Invalid refresh signed: %w", err)
	}

	accessToken.TokenString = accessTokenSigned
	refreshToken.TokenString = refreshTokenSigned

	return accessToken, refreshToken, nil
}
