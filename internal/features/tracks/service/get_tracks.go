package tracks_service

import (
	"context"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
	"github.com/google/uuid"
)

func (s *TrackServise) GetTracks(ctx context.Context, trackId uuid.UUID) ([]domain.Track, error) {
	return []domain.Track{}, nil
}
