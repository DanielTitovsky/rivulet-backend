package users_transport_http

import (
	"context"
	"net/http"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
	app_http_server "github.com/DanielTitovsky/rivulet-backend.git/internal/app/transport/http/server"
	"github.com/google/uuid"
)

type UsersHttpHandler struct {
	userService UserService
}

type UserService interface {
	CreateUser(ctx context.Context, user domain.User) (domain.User, error)
	GetUser(ctx context.Context, userId uuid.UUID) (domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (domain.User, error)
	DeleteUser(ctx context.Context, UserId uuid.UUID) error
	UpdateUser(ctx context.Context, userId uuid.UUID, updateUser domain.UserUpdate) (domain.User, error)
	AddTrackToFavorite(ctx context.Context, userId uuid.UUID, trackId uuid.UUID) error
	RemoveTrackFromFavorite(ctx context.Context, userId uuid.UUID, trackId uuid.UUID) error
	GetOrCreateOAuthUser(ctx context.Context, email string, name string) (domain.User, error)
}

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
			Handler: h.CreateUser,
		},
		{
			Method:  http.MethodGet,
			Path:    "/user/:id",
			Handler: h.GetUser,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/user/:id",
			Handler: h.DeleteUser,
		},
		{
			Method:  http.MethodPatch,
			Path:    "/user/:id",
			Handler: h.UpdateUser,
		},
		{
			Method:  http.MethodPost,
			Path:    "/user/:id/favorite-track/:trackId",
			Handler: h.AddTrackToFavorite,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/user/:id/favorite-track/:trackId",
			Handler: h.RemoveTrackFromFavorite,
		},
	}
}
