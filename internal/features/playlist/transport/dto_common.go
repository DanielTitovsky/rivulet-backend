package playlist_transport_http

import (
	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
	"github.com/google/uuid"
)

type PlaylistDTOResponce struct {
	Id          uuid.UUID      `json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	CoverUrl    string         `json:"coverUrl"`
	Tracks      []domain.Track `json:"tracks"`
	Visibility  bool           `json:"visibility"`
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
