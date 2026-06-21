package playlist_service

import (
	"context"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
	"github.com/google/uuid"
)

type PlaylistService struct {
}

type PlayListRepository interface {
	GetPlaylist(ctx context.Context, playlistId uuid.UUID) (domain.Playlsit, error)
	GetPlaylistUser(ctx context.Context, playlistId uuid.UUID) (domain.Playlsit, error)
}

type PlayListStorage interface {
	GetPlayList(ctx context.Context, playlistId uuid.UUID) (domain.Playlsit, error)
}
