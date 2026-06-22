package artist_transport_http

import (
	"time"

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
type ArtistAlbumDTOResponse struct {
	Id           uuid.UUID
	Name         string
	Description  string
	CoverUrl     string
	Release_date time.Time
}

func ArtistAlbumDTOFromDomain(album domain.Album) ArtistAlbumDTOResponse {
	return ArtistAlbumDTOResponse{
		Id:           album.Id,
		Name:         album.Name,
		Description:  album.Description,
		CoverUrl:     album.CoverUrl,
		Release_date: album.Release_date,
	}
}

func ArtistAlbumsDTOFromDomain(albums []domain.Album) []ArtistAlbumDTOResponse {
	response := make([]ArtistAlbumDTOResponse, 0, len(albums))

	for _, album := range albums {
		response = append(response, ArtistAlbumDTOFromDomain(album))
	}

	return response
}

func ArtistDTOFromDomain(artist domain.Artist) ArtistDTOResponse {
	return ArtistDTOResponse{
		Id:          artist.Id,
		Name:        artist.Name,
		AvatarUrl:   artist.AvatarUrl,
		Description: artist.Description,
	}
}
