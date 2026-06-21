package auth_transport_http

import (
	"context"
	"net/http"
	"time"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
	app_oauth "github.com/DanielTitovsky/rivulet-backend.git/internal/app/oauth"
	app_http_server "github.com/DanielTitovsky/rivulet-backend.git/internal/app/transport/http/server"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type AuthHttpHandler struct {
	authService       AuthService
	googleOAuthConfig *oauth2.Config
	frontendURL       string
}

const (
	AccessExpired  = 15 * time.Minute
	RefreshExpired = 30 * 24 * time.Hour
)

type AuthService interface {
	Register(ctx context.Context, user domain.User) (domain.User, *domain.Token, *domain.Token, error)
	Login(ctx context.Context, email string, password string) (domain.User, *domain.Token, *domain.Token, error)
	Refresh(ctx context.Context, refreshToken string) (domain.User, *domain.Token, *domain.Token, error)
	Logout(ctx context.Context, refreshToken string) error
	Me(ctx context.Context, accessToken string) (domain.User, error)
	OAuthLogin(ctx context.Context, oauthUser domain.OAuthUser) (domain.User, *domain.Token, *domain.Token, error)
}

func NewAuthHttpHandler(
	authService AuthService,
	oauthConfig app_oauth.OAuthConfig,
) *AuthHttpHandler {
	return &AuthHttpHandler{
		authService: authService,
		googleOAuthConfig: &oauth2.Config{
			ClientID:     oauthConfig.GoogleClientID,
			ClientSecret: oauthConfig.GoogleClientSecret,
			RedirectURL:  oauthConfig.GoogleRedirectURL,
			Scopes: []string{
				"openid",
				"email",
				"profile",
			},
			Endpoint: google.Endpoint,
		},
		frontendURL: oauthConfig.FrontendURL,
	}
}

func (h *AuthHttpHandler) Routers() []app_http_server.Route {
	return []app_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/register/",
			Handler: h.Register,
		},
		{
			Method:  http.MethodPost,
			Path:    "/login/",
			Handler: h.Login,
		},
		{
			Method:  http.MethodPost,
			Path:    "/refresh/",
			Handler: h.Refresh,
		},
		{
			Method:  http.MethodPost,
			Path:    "/logout/",
			Handler: h.Logout,
		},
		{
			Method:  http.MethodGet,
			Path:    "/me/",
			Handler: h.Me,
		},
		{
			Method:  http.MethodGet,
			Path:    "/google/",
			Handler: h.GoogleLogin,
		},
		{
			Method:  http.MethodGet,
			Path:    "/google/callback/",
			Handler: h.GoogleCallback,
		},
	}
}
