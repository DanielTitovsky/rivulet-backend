package token_transport_http

import (
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
	RemoveToken(tokenId uuid.UUID) error
	GetToken(tokenString string) (*domain.Token, error)
	RefreshToken(tokenString string) (*domain.Token, error)
	ValidateToken(cookieToken *http.Cookie, tokenType string) (*domain.TokenClaims, error)
	SaveToken(token domain.Token) (*domain.Token, error)
	GenerateTokens(accessExpires time.Duration, refreshExpires time.Duration, user domain.User) (*domain.Token, *domain.Token, error)
}

func (th *TokenHttpHandler) Routers() []app_http_server.Route {
	return []app_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/refresh/",
			Handler: th.RefreshToken,
		},
	}
}
