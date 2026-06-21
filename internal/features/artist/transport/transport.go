package artist_transport_http

import (
	"context"
	"net/http"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
	app_http_server "github.com/DanielTitovsky/rivulet-backend.git/internal/app/transport/http/server"
	"github.com/google/uuid"
)

type ArtistHttpHandler struct {
	ArtistService
}

type ArtistService interface {
	GetArtist(ctx context.Context, artistId uuid.UUID) (domain.Artist, error)
}

func NewArtistHttpHandler(artistService ArtistService) *ArtistHttpHandler {
	return &ArtistHttpHandler{
		ArtistService: artistService,
	}
}

func (h *ArtistHttpHandler) Routers() []app_http_server.Route {
	return []app_http_server.Route{
		{
			Method:  http.MethodGet,
			Path:    "/track/",
			Handler: h.getArtist,
		},
	}
}
