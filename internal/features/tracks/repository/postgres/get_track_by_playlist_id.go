package tracks_postgres_repository

import (
	"context"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
	"github.com/google/uuid"
)

func (r *TrackRepository) GetTracksByPlaylistId(ctx context.Context, playlistId uuid.UUID) ([]domain.Track, error) {

	return []domain.Track{}, nil
}
