package tracks_service

import (
	"context"
	"fmt"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
	"github.com/google/uuid"
)

func (s *TrackServise) GetTrack(ctx context.Context, trackId uuid.UUID) (domain.Track, error) {
	track, err := s.TrackRepository.GetTrack(ctx, trackId)

	if err != nil {
		return domain.Track{}, fmt.Errorf("Get track by id: %w", err)
	}

	return track, nil
}
