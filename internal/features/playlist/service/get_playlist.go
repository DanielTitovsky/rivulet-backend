package playlist_service

import (
	"context"
	"fmt"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
	"github.com/google/uuid"
)

func (s *PlaylistService) GetPlaylist(ctx context.Context, playlistId uuid.UUID) (domain.Playlsit, error) {
	playlist, err := s.PlaylistRepository.GetPlaylist(ctx, playlistId)

	if err != nil {
		return domain.Playlsit{}, fmt.Errorf("Get playlist by id: %w", err)
	}

	tracks, err := s.TrackService.GetTracksByPlaylistId(ctx, playlistId)

	if err != nil {
		return domain.Playlsit{}, fmt.Errorf("Get playlist tracks: %w", err)
	}

	fmt.Print("\n")
	fmt.Print("\n")
	fmt.Print(tracks)
	fmt.Print("\n")
	fmt.Print("\n")

	playlist.Tracks = tracks

	playlist.CoverUrl, err = s.PlaylisttorageRepository.GetPlaylistCoverLink(ctx, playlist.CoverUrl)

	if err != nil {
		return domain.Playlsit{}, fmt.Errorf("Get playlist cover: %w", err)
	}

	return playlist, nil
}
