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

	for index, track := range tracks {
		if track.AudioStorageKey != "" {
			trackAudioLink, err := s.TrackStorageRepository.GetTrackAudioLink(ctx, track.AudioStorageKey)

			if err != nil {
				return nil, fmt.Errorf("Get track audio link: %w", err)
			}

			tracks[index].AudioStorageKey = trackAudioLink
		}

		if track.CoverStorageKey != "" {
			trackCoverLink, err := s.TrackStorageRepository.GetTrackCoverLink(ctx, track.CoverStorageKey)

			if err != nil {
				return nil, fmt.Errorf("Get track cover link: %w", err)
			}

			tracks[index].CoverStorageKey = trackCoverLink
		}
	}

	return tracks, nil
}
