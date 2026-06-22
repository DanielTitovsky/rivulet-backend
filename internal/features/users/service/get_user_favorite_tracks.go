package users_service

import (
	"context"
	"fmt"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
	"github.com/google/uuid"
)

func (s *UsersServise) GetUserFavoriteTracks(
	ctx context.Context,
	userId uuid.UUID,
) ([]domain.Track, error) {
	tracks, err := s.UsersRepository.GetUserFavoriteTracks(ctx, userId)

	if err != nil {
		return nil, fmt.Errorf("Get user favorite tracks: %w", err)
	}

	for index, track := range tracks {
		if track.AudioStorageKey != "" {
			trackAudioLink, err := s.UserStorageRepository.GetTrackAudioLink(ctx, track.AudioStorageKey)

			if err != nil {
				return nil, fmt.Errorf("Get track audio link: %w", err)
			}

			tracks[index].AudioStorageKey = trackAudioLink
		}

		if track.CoverStorageKey != "" {
			trackCoverLink, err := s.UserStorageRepository.GetTrackCoverLink(ctx, track.CoverStorageKey)

			if err != nil {
				return nil, fmt.Errorf("Get track cover link: %w", err)
			}

			tracks[index].CoverStorageKey = trackCoverLink
		}
	}

	return tracks, nil
}
