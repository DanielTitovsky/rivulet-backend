package token_service

import (
	"fmt"
	"net/http"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
	"github.com/golang-jwt/jwt"
)

func (th *TokenService) ValidateToken(cookieToken *http.Cookie, tokenType string) (*domain.TokenClaims, error) {

	claims := &domain.TokenClaims{}

	verifiedToken, err := jwt.ParseWithClaims(cookieToken.Value, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		switch tokenType {
		case "access":
			return []byte(th.tokenAccessSecret), nil
		case "refresh":
			return []byte(th.tokenRefreshSecret), nil
		default:
			return nil, fmt.Errorf("Invalid token type")
		}

	})

	if err != nil {
		return nil, err
	}

	if !verifiedToken.Valid {
		return nil, fmt.Errorf("Failed validate %s token: %w", tokenType, err)
	}

	return claims, nil
}
