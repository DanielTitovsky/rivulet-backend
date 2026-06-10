package auth_transport_http

import (
	"context"
	"net/http"
	"time"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
	app_http_server "github.com/DanielTitovsky/rivulet-backend.git/internal/app/transport/http/server"
)

type AuthHttpHandler struct {
	authService AuthService
}

type AuthService interface {
	RegisterUser(ctx context.Context, user domain.User) (domain.User, error)
	LoginUser(ctx context.Context, user domain.User) error
	LogoutUser(ctx context.Context, user domain.User) error
	RegisterGoogleUser(ctx context.Context) error
	RegisterGoogleUserCallback(ctx context.Context, googleCode string) error
	GenerateTokens(ctx context.Context, AccessTokenExpires time.Duration, RefreshTokenExpires time.Duration, userDomain domain.User) (domain.Token, domain.Token, error)
	SaveToken(ctx context.Context, RefreshToken domain.Token) (domain.Token, error)
}

func NewAuthHttpHandler(authService AuthService) *AuthHttpHandler {
	return &AuthHttpHandler{
		authService: authService,
	}
}

func (h *AuthHttpHandler) Routers() []app_http_server.Route {
	return []app_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/register/",
			Handler: h.RegisterUser,
		},
		{
			Method:  http.MethodPost,
			Path:    "/login/",
			Handler: h.LoginUser,
		},
		{
			Method:  http.MethodPost,
			Path:    "/logout/",
			Handler: h.LogoutUser,
		},
		{
			Method:  http.MethodPost,
			Path:    "/register/google/",
			Handler: h.RegisterGoogleUser,
		},
		{
			Method:  http.MethodPost,
			Path:    "/register/google/callback/",
			Handler: h.RegisterGoogleUserCallback,
		},
	}
}
