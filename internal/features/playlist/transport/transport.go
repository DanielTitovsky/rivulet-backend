package playlist_transport_http

import (
	"context"
	"net/http"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
	app_http_server "github.com/DanielTitovsky/rivulet-backend.git/internal/app/transport/http/server"
	"github.com/google/uuid"
)

type PlaylistHttpHandler struct {
	PlaylistService
}

type PlaylistService interface {
	GetPlaylist(ctx context.Context, playlistId uuid.UUID) (domain.Playlsit, error)
	GetUserPlaylists(ctx context.Context, userId uuid.UUID) ([]domain.Playlsit, error)
	GetPlaylists(ctx context.Context) ([]domain.Playlsit, error)
}

func NewPlaylistHttpHandler(playlistService PlaylistService) *PlaylistHttpHandler {
	return &PlaylistHttpHandler{
		PlaylistService: playlistService,
	}
}

func (h *PlaylistHttpHandler) Routers() []app_http_server.Route {
	return []app_http_server.Route{
		{
			Method:  http.MethodGet,
			Path:    "/playlist/:id",
			Handler: h.GetPlaylist,
		},
		{
			Method:  http.MethodGet,
			Path:    "/user/:id/playlist/",
			Handler: h.GetUserPlaylists,
		},
		{
			Method:  http.MethodGet,
			Path:    "/playlist/",
			Handler: h.GetPlaylists,
		},
	}
}
