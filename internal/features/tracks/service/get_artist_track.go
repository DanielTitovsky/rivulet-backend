package tracks_service

import (
	"context"
	"fmt"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
	"github.com/google/uuid"
)

func (s *TrackServise) GetTracksByArtistId(
	ctx context.Context,
	artistId uuid.UUID,
) ([]domain.Track, error) {
	tracks, err := s.TrackRepository.GetTracksByArtistId(ctx, artistId)

	if err != nil {
		return nil, fmt.Errorf("get tracks by artist id: %w", err)
	}

	tracks, err = s.attachTrackLinks(ctx, tracks)
	if err != nil {
		return nil, fmt.Errorf("attach track links: %w", err)
	}

	return tracks, nil
}

func (s *TrackServise) attachTrackLinks(
	ctx context.Context,
	tracks []domain.Track,
) ([]domain.Track, error) {
	for index, track := range tracks {
		trackAudioLink, err := s.TrackStorageRepository.GetTrackAudioLink(
			ctx,
			track.AudioStorageKey,
		)
		if err != nil {
			return nil, fmt.Errorf("get track audio link: %w", err)
		}

		trackCoverLink, err := s.TrackStorageRepository.GetTrackCoverLink(
			ctx,
			track.CoverStorageKey,
		)
		if err != nil {
			return nil, fmt.Errorf("get track cover link: %w", err)
		}

		tracks[index].AudioStorageKey = trackAudioLink
		tracks[index].CoverStorageKey = trackCoverLink
	}

	return tracks, nil
}
