package token_transport_http

import (
	"context"
	"net/http"
	"time"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
	app_http_server "github.com/DanielTitovsky/rivulet-backend.git/internal/app/transport/http/server"
	uuid "github.com/google/uuid"
)

type TokenHttpHandler struct {
	tokenService TokenService
}

type TokenService interface {
	RemoveToken(ctx context.Context, tokenId uuid.UUID) error
	GetRefreshToken(ctx context.Context, rawToken string) (domain.Token, error)
	ValidateToken(tokenString string, tokenType string) (*domain.TokenClaims, error)
	SaveRefreshToken(ctx context.Context, token domain.Token) (domain.Token, error)
	GenerateTokens(accessExpires time.Duration, refreshExpires time.Duration, user domain.User) (*domain.Token, *domain.Token, error)
}

func NewTokensHttpHandler(tokenService TokenService) *TokenHttpHandler {
	return &TokenHttpHandler{
		tokenService: tokenService,
	}
}

func (th *TokenHttpHandler) Routers() []app_http_server.Route {
	return []app_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/fasdfasd/",
			Handler: th.RefreshToken,
		},
	}
}
