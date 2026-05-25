package users_transport_http

import (
	"context"
	"net/http"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
	app_http_server "github.com/DanielTitovsky/rivulet-backend.git/internal/app/transport/http/server"
)

type UsersHttpHandler struct {
	userService UserService
}

type UserService interface {
	CreateUser(ctx context.Context, user domain.User) (domain.User, error)
}

// GetUser(userIdentity string) (*domain.User, error)
// 	UpdateUser(updateUser domain.User) (domain.User, error)
// 	CreateProviderUser(provideUser domain.ProvideUser, userId uuid.UUID) error
// 	DeleteUser(UserId uuid.UUID) bool
// 	Login(unverifiedUser *domain.User, pasword string) (*domain.User, error)

func NewUsersHttpHandler(userService UserService) *UsersHttpHandler {
	return &UsersHttpHandler{
		userService: userService,
	}
}

func (h *UsersHttpHandler) Routers() []app_http_server.Route {
	return []app_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/user/",
			Handler: h.RegisterUser,
		},
	}
}
