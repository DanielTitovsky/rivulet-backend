package tracks_transport_http

import (
	"time"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
	"github.com/google/uuid"
)

type TrackDTOResponse struct {
	Id              uuid.UUID
	Name            string
	Description     string
	Cover           string
	AudioStorageKey string
	AlbumId         uuid.UUID
	Artists         []domain.Artist
	Duration        time.Duration
	IsExplicit      bool
	IsStreamable    bool
	IsDownloadable  bool
	IsFavorite      bool
	Status          string
}

func trackDTOFromDomain(track domain.Track) TrackDTOResponse {
	return TrackDTOResponse{
		Id:              track.Id,
		Name:            track.Name,
		Description:     track.Description,
		Cover:           track.CoverStorageKey,
		AudioStorageKey: track.AudioStorageKey,
		AlbumId:         *track.AlbumId,
		Artists:         track.Artists,
		Duration:        track.DurationSecond,
		IsExplicit:      track.IsExplicit,
		IsStreamable:    track.IsStreamable,
		IsDownloadable:  track.IsDownloadable,
		Status:          track.Status,
	}
}

func tracksDTOFromDomain(track []domain.Track) []TrackDTOResponse {
	tracksDTO := make([]TrackDTOResponse, len(track))

	for _, domainTrack := range track {
		tracksDTO = append(tracksDTO, trackDTOFromDomain(domainTrack))
	}

	return tracksDTO
}
