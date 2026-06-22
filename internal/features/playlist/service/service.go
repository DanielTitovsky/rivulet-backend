package playlist_service

import (
	"context"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
	"github.com/google/uuid"
)

type PlaylistService struct {
	PlaylistRepository       PlaylistRepository
	PlaylisttorageRepository PlaylisttorageRepository
	TrackService             TrackService
}

type PlaylistRepository interface {
	GetPlaylist(ctx context.Context, playlistId uuid.UUID) (domain.Playlsit, error)
	GetUserPlaylists(ctx context.Context, userId uuid.UUID) ([]domain.Playlsit, error)
	GetPlaylists(ctx context.Context) ([]domain.Playlsit, error)
}

type PlaylisttorageRepository interface {
	GetPlaylistCoverLink(ctx context.Context, link string) (string, error)
}

type TrackService interface {
	GetTracksByPlaylistId(ctx context.Context, playlistId uuid.UUID) ([]domain.Track, error)
}

func NewPlaylistService(
	playlistRepository PlaylistRepository,
	trackService TrackService,
	playlisttorageRepository PlaylisttorageRepository,
) *PlaylistService {
	return &PlaylistService{
		PlaylistRepository:       playlistRepository,
		TrackService:             trackService,
		PlaylisttorageRepository: playlisttorageRepository,
	}
}
