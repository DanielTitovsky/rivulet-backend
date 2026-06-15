package tracks_transport_http

import (
	"context"
	"net/http"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
	app_http_server "github.com/DanielTitovsky/rivulet-backend.git/internal/app/transport/http/server"
	"github.com/google/uuid"
)

type TrackHttpHandler struct {
	trackService TrackService
}

type TrackService interface {
	CreateTrack(ctx context.Context, track domain.Track) (domain.Track, error)
	GetTrack(ctx context.Context, trackId uuid.UUID) (domain.Track, error)
	GetTracks(ctx context.Context, trackId uuid.UUID) ([]domain.Track, error)
	DeleteTrack(ctx context.Context, trackId uuid.UUID) error
	UpdateTrack(ctx context.Context, trackId uuid.UUID, updateTrack domain.TrackUpdate) (domain.Track, error)
}

func NewTrackHttpHandler(trackService TrackService) *TrackHttpHandler {
	return &TrackHttpHandler{
		trackService: trackService,
	}
}

func (h *TrackHttpHandler) Routers() []app_http_server.Route {
	return []app_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/track/",
			Handler: h.CreateTrack,
		},
		{
			Method:  http.MethodGet,
			Path:    "/track/:id",
			Handler: h.GetTrack,
		},
		{
			Method:  http.MethodGet,
			Path:    "/tracks/",
			Handler: h.GetTracks,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/track/:id",
			Handler: h.DeleteTrack,
		},
		{
			Method:  http.MethodPatch,
			Path:    "/track/:id",
			Handler: h.UpdateTrack,
		},
	}
}
