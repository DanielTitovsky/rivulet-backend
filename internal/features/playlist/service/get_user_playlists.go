package playlist_service

import (
	"context"
	"fmt"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
	"github.com/google/uuid"
)

func (s *PlaylistService) GetUserPlaylists(ctx context.Context, userId uuid.UUID) ([]domain.Playlsit, error) {
	playlists, err := s.PlaylistRepository.GetUserPlaylists(ctx, userId)

	if err != nil {
		return nil, fmt.Errorf("Get user playlists: %w", err)
	}

	for i := range playlists {
		if playlists[i].CoverUrl == "" {
			continue
		}

		playlists[i].CoverUrl, err = s.PlaylisttorageRepository.GetPlaylistCoverLink(ctx, playlists[i].CoverUrl)

		if err != nil {
			return nil, fmt.Errorf("Get playlist cover: %w", err)
		}
	}

	return playlists, nil
}
