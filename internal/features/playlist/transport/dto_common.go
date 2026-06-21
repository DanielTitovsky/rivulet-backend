package playlist_transport_http

import (
	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
	"github.com/google/uuid"
)

type PlaylistDTOResponce struct {
	Id          uuid.UUID
	Name        string
	Description string
	CoverUrl    string
	Tracks      []domain.Track
	Visibility  bool
}

func playlistDTOFromDomain(playlist domain.Playlsit) PlaylistDTOResponce {
	return PlaylistDTOResponce{
		Id:          playlist.Id,
		Name:        playlist.Name,
		Description: playlist.Description,
		CoverUrl:    playlist.CoverUrl,
		Tracks:      playlist.Tracks,
		Visibility:  playlist.Visibility,
	}
}
