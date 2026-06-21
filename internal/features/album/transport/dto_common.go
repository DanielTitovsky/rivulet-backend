package album_transport_http

import (
	"time"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
	"github.com/google/uuid"
)

type AlbumDTOResponse struct {
	Id           uuid.UUID
	Artists      []domain.Artist
	Name         string
	CoverUrl     string
	Description  string
	Release_date time.Time
}

func albumDTOFromDomain(album domain.Album) AlbumDTOResponse {
	return AlbumDTOResponse{
		Id:           album.Id,
		Artists:      album.Artists,
		Name:         album.Name,
		CoverUrl:     album.CoverUrl,
		Description:  album.Description,
		Release_date: album.Release_date,
	}
}

func albumsDTOFromDomain(albums []domain.Album) []AlbumDTOResponse {
	albumsDTO := make([]AlbumDTOResponse, 0, len(albums))

	for _, domainAlbum := range albums {
		albumsDTO = append(albumsDTO, albumDTOFromDomain(domainAlbum))
	}

	return albumsDTO
}
