package artist_transport_http

import (
	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
	"github.com/google/uuid"
)

type ArtistDTOResponse struct {
	Id               uuid.UUID
	Name             string
	AvatarUrl        string
	Description      string
	AvatarStorageKey string
}

func ArtistDTOFromDomain(artist domain.Artist) ArtistDTOResponse {
	return ArtistDTOResponse{
		Id:          artist.Id,
		Name:        artist.Name,
		AvatarUrl:   artist.AvatarUrl,
		Description: artist.Description,
	}
}
