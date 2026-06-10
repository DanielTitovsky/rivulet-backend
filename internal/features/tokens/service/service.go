package token_service

import (
	"context"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
	"github.com/google/uuid"
)

type TokenService struct {
	TokenRepository    TokenRepository
	tokenRefreshSecret string
	tokenAccessSecret  string
}

type TokenRepository interface {
	SaveToken(ctx context.Context, token *domain.Token) (*domain.Token, error)
	GetToken(ctx context.Context, tokenId uuid.UUID) (*domain.Token, error)
	UpdateUserToken(ctx context.Context, token domain.Token) (*domain.Token, error)
	RemoveToken(ctx context.Context, tokenId uuid.UUID) error
}

func NewTokenService(rep TokenRepository, tokenRefreshSecret string, tokenAccessSecret string) *TokenService {
	return &TokenService{
		TokenRepository:    rep,
		tokenRefreshSecret: tokenRefreshSecret,
		tokenAccessSecret:  tokenAccessSecret,
	}
}
