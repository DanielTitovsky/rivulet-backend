package album_transport_http

import (
	"context"
	"net/http"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
	app_http_server "github.com/DanielTitovsky/rivulet-backend.git/internal/app/transport/http/server"
	"github.com/google/uuid"
)

type AlbumHttpHandler struct {
	AlbumService
}

type AlbumService interface {
	GetAlbum(ctx context.Context, albumId uuid.UUID) (domain.Album, error)
	GetArtistAbums(ctx context.Context, artistId uuid.UUID) ([]domain.Album, error)
}

func NewAlbumHttpHandler(albumService AlbumService) *AlbumHttpHandler {
	return &AlbumHttpHandler{
		AlbumService: albumService,
	}
}

func (h *AlbumHttpHandler) Routers() []app_http_server.Route {
	return []app_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/track/",
			Handler: h.GetAlbum,
		},
		{
			Method:  http.MethodGet,
			Path:    "/track/:id",
			Handler: h.GetArtistAbums,
		},
	}
}
