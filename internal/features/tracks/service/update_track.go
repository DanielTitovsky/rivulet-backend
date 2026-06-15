package tracks_service

import (
	"context"
	"fmt"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
	"github.com/google/uuid"
)

func (s *TrackServise) UpdateTrack(ctx context.Context, trackId uuid.UUID, updateTrack domain.TrackUpdate) (domain.Track, error) {
	if err := updateTrack.Validate(); err != nil {
		return domain.Track{}, fmt.Errorf("Validate updateUser: %w", err)
	}

	track, err := s.TrackRepository.GetTrack(ctx, trackId)

	if err != nil {
		return domain.Track{}, fmt.Errorf("Get track: %w", err)
	}

	if err := track.ApplyUpdate(updateTrack); err != nil {
		return domain.Track{}, fmt.Errorf("apply track update: %w", err)
	}

	err = s.startUpdateTrackTx(ctx, trackId, track)

	if err != nil {
		return domain.Track{}, fmt.Errorf("Failed to update track: %w", err)
	}

	updatedTrack, err := s.GetTrack(ctx, trackId)

	if err != nil {
		return domain.Track{}, fmt.Errorf("Failed to update track: %w", err)
	}

	return updatedTrack, nil
}

func (s *TrackServise) startUpdateTrackTx(ctx context.Context, trackId uuid.UUID, track domain.Track) error {
	return s.TransactionManager.WithinTransaction(ctx, func(ctx context.Context) error {
		err := s.TrackRepository.UpdateTrack(ctx, trackId, track)

		if err != nil {
			return fmt.Errorf("Failed update track: %w", err)
		}

		err = s.TrackRepository.UpdateTrackArtists(ctx, trackId, track.ArtistIds)

		if err != nil {
			return fmt.Errorf("Failed update track_artists: %w", err)
		}

		err = s.TrackRepository.UpdateTrackGenres(ctx, trackId, track.GenreIds)

		if err != nil {
			return fmt.Errorf("Failed update track_genres: %w", err)
		}

		return nil
	})
}
