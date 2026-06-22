package artist_service

import (
	"context"
	"fmt"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
	"github.com/google/uuid"
)

func (s *ArtistService) GetArtistAlbums(
	ctx context.Context,
	artistId uuid.UUID,
) ([]domain.Album, error) {
	albums, err := s.ArtistRepository.GetArtistAlbums(ctx, artistId)

	if err != nil {
		return nil, fmt.Errorf("Select artist albums: %w", err)
	}

	for i := range albums {
		if albums[i].CoverUrl == "" {
			continue
		}

		albums[i].CoverUrl, err = s.ArtistStorage.GetArtistAvatar(ctx, albums[i].CoverUrl)

		if err != nil {
			return nil, fmt.Errorf("Get playlist cover: %w", err)
		}
	}

	return albums, nil
}
