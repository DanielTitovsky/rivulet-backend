package tokens_service

import (
	"fmt"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
	"github.com/golang-jwt/jwt"
)

func (s *TokensServise) ValidateToken(
	tokenString string,
	tokenType string,
) (*domain.TokenClaims, error) {
	claims := &domain.TokenClaims{}

	verifiedToken, err := jwt.ParseWithClaims(
		tokenString,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			switch tokenType {
			case "access":
				return s.accessSecret, nil
			case "refresh":
				return s.refreshSecret, nil
			default:
				return nil, fmt.Errorf("invalid token type")
			}
		},
	)

	if err != nil {
		return nil, fmt.Errorf("parse token: %w", err)
	}

	if !verifiedToken.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
