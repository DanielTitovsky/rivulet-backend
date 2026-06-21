package tokens_service

import (
	"context"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
	"github.com/google/uuid"
)

type TokenRepository interface {
	SaveToken(ctx context.Context, token domain.Token) (domain.Token, error)
	GetTokenByHash(ctx context.Context, tokenHash string) (domain.Token, error)
	RemoveToken(ctx context.Context, tokenId uuid.UUID) error
}

type TokensServise struct {
	TokenRepository TokenRepository
	accessSecret    []byte
	refreshSecret   []byte
}

func NewTokensServise(
	tokenRepository TokenRepository,
	accessSecret string,
	refreshSecret string,
) *TokensServise {
	return &TokensServise{
		TokenRepository: tokenRepository,
		accessSecret:    []byte(accessSecret),
		refreshSecret:   []byte(refreshSecret),
	}
}
