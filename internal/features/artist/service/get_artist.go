package artist_service

import (
	"context"
	"fmt"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
	"github.com/google/uuid"
)

func (s *ArtistService) GetArtist(ctx context.Context, artistId uuid.UUID) (domain.Artist, error) {
	artist, err := s.ArtistRepository.GetArtist(ctx, artistId)

	if err != nil {
		return domain.Artist{}, fmt.Errorf("Select artist: %w", err)
	}

	artist.AvatarUrl, err = s.ArtistStorage.GetArtistAvatar(ctx, artistId)

	if err != nil {
		return domain.Artist{}, fmt.Errorf("Get artist avatar: %w", err)
	}

	return artist, nil
}
