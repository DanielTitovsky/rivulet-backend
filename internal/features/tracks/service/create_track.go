package tracks_service

import (
	"context"
	"fmt"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
	"github.com/google/uuid"
)

func (s *TrackServise) CreateTrack(ctx context.Context, track domain.Track) (domain.Track, error) {
	if err := track.Validate(); err != nil {
		return domain.Track{}, fmt.Errorf("Validate track domain: %w", err)
	}

	track, err := s.startCreateTrackTx(ctx, track)

	if err != nil {
		fmt.Print("\n")
		fmt.Print(err)
		fmt.Print("\n")
		return domain.Track{}, fmt.Errorf("Failed to create track: %w", err)
	}

	return track, nil
}

func (s *TrackServise) startCreateTrackTx(ctx context.Context, track domain.Track) (domain.Track, error) {
	var trackId uuid.UUID

	err := s.TransactionManager.WithinTransaction(ctx, func(ctx context.Context) error {
		var err error

		trackId, err = s.TrackRepository.CreateTrack(ctx, track)

		if err != nil {
			return fmt.Errorf("Failed create track: %w", err)
		}

		err = s.TrackRepository.CreateTrackArtists(ctx, trackId, track.ArtistIds)

		if err != nil {
			return fmt.Errorf("Failed create track_artists: %w", err)
		}

		err = s.TrackRepository.CreateTrackGenres(ctx, trackId, track.GenreIds)

		if err != nil {
			return fmt.Errorf("Failed create track_genres: %w", err)
		}

		return nil
	})

	if err != nil {
		return domain.Track{}, fmt.Errorf("Failed to create track: %w", err)
	}

	createdTrack, err := s.TrackRepository.GetTrack(ctx, trackId)

	if err != nil {
		return domain.Track{}, fmt.Errorf("Failed to get new track: %w", err)
	}

	return createdTrack, nil
}
