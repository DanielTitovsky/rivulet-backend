package playlist_service

import (
	"context"
	"fmt"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
)

func (s *PlaylistService) GetPlaylists(ctx context.Context) ([]domain.Playlsit, error) {
	playlists, err := s.PlaylistRepository.GetPlaylists(ctx)

	if err != nil {
		return nil, fmt.Errorf("Get playlists: %w", err)
	}

	for index, playlist := range playlists {
		tracks, err := s.TrackService.GetTracksByPlaylistId(ctx, playlist.Id)

		if err != nil {
			return nil, fmt.Errorf("Get playlist tracks: %w", err)
		}

		playlists[index].Tracks = tracks

		if playlist.CoverUrl == "" {
			continue
		}

		playlists[index].CoverUrl, err = s.PlaylisttorageRepository.GetPlaylistCoverLink(ctx, playlist.CoverUrl)

		if err != nil {
			return nil, fmt.Errorf("Get playlist cover: %w", err)
		}
	}

	return playlists, nil
}
