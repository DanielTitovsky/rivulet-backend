package tracks_service

import (
	"context"
	"fmt"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
)

func (s *TrackServise) GetRandomTrack(ctx context.Context) (domain.Track, error) {
	track, err := s.TrackRepository.GetRandomTrack(ctx)

	if err != nil {
		return domain.Track{}, fmt.Errorf("Get track by id: %w", err)
	}

	trackAudioLink, err := s.TrackStorageRepository.GetTrackAudioLink(ctx, track.AudioStorageKey)
	if err != nil {
		return domain.Track{}, fmt.Errorf("get track audio link: %w", err)
	}

	trackCoverLink, err := s.TrackStorageRepository.GetTrackCoverLink(ctx, track.CoverStorageKey)
	if err != nil {
		return domain.Track{}, fmt.Errorf("get track cover link: %w", err)
	}

	track.AudioStorageKey = trackAudioLink
	track.CoverStorageKey = trackCoverLink

	return track, nil
}
