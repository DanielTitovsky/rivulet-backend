package tracks_service

import (
	"context"
	"fmt"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
	"github.com/google/uuid"
)

func (s *TrackServise) GetTracksByPlaylistId(
	ctx context.Context,
	playlistId uuid.UUID,
) ([]domain.Track, error) {
	tracks, err := s.TrackRepository.GetTracksByPlaylistId(ctx, playlistId)

	if err != nil {
		return nil, fmt.Errorf("get tracks by playlist id: %w", err)
	}

	tracks, err = s.attachTrackLinks(ctx, tracks)

	fmt.Print("\n")
	fmt.Print("\n")
	fmt.Print(tracks)
	fmt.Print("\n")
	fmt.Print("\n")

	if err != nil {
		return nil, fmt.Errorf("attach track links: %w", err)
	}

	return tracks, nil
}
