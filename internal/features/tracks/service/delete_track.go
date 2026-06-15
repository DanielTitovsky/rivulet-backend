package tracks_service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (s *TrackServise) DeleteTrack(ctx context.Context, trackId uuid.UUID) error {
	err := s.TrackRepository.DeleteTrack(ctx, trackId)

	if err != nil {
		return fmt.Errorf("Delete track: %w", err)
	}

	return nil
}
