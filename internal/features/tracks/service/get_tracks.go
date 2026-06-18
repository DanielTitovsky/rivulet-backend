package tracks_service

import (
	"context"
	"fmt"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
)

func (s *TrackServise) GetTracks(ctx context.Context, trackFilter domain.TrackFilters) ([]domain.Track, error) {
	if trackFilter.Limit <= 0 {
		return nil, fmt.Errorf("limit must be non-negative")
	}

	if trackFilter.Offset < 0 {
		return nil, fmt.Errorf("ofset must be non-negative")
	}

	tracks, err := s.TrackRepository.GetTracks(ctx, trackFilter)

	if err != nil {
		return nil, fmt.Errorf("Failed to get tracks: %w", err)
	}

	return tracks, nil
}
