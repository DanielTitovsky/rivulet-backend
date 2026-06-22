package users_transport_http

import (
	"time"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
	"github.com/google/uuid"
)

type UserDTOResponse struct {
	Id    uuid.UUID `json:"id" validate:"required"`
	Name  string    `json:"name" validate:"required"`
	Email string    `json:"email"`
}

func userDTOFromDomain(user domain.User) UserDTOResponse {
	return UserDTOResponse{
		Id:    user.Id,
		Name:  user.Name,
		Email: *user.Email,
	}
}

type TrackDTOResponse struct {
	Id              uuid.UUID
	Name            string
	Description     string
	AlbumId         *uuid.UUID
	ArtistIds       []uuid.UUID
	Artists         []TrackArtistDTOResponse
	GenreIds        []uuid.UUID
	Genres          []string
	CoverStorageKey string
	DurationSecond  int
	ReleaseDate     time.Time
	IsExplicit      bool
	IsStreamable    bool
	IsDownloadable  bool
	StatusId        int
	Status          string
	AudioStorageKey string
}

type TrackArtistDTOResponse struct {
	Id          uuid.UUID
	Name        string
	AvatarUrl   string
	Description string
}

func tracksDTOFromDomain(tracks []domain.Track) []TrackDTOResponse {
	response := make([]TrackDTOResponse, 0, len(tracks))

	for _, track := range tracks {
		response = append(response, trackDTOFromDomain(track))
	}

	return response
}

func trackDTOFromDomain(track domain.Track) TrackDTOResponse {
	artistIds := make([]uuid.UUID, 0, len(track.Artists))
	artists := make([]TrackArtistDTOResponse, 0, len(track.Artists))

	for _, artist := range track.Artists {
		artistIds = append(artistIds, artist.Id)

		artists = append(artists, TrackArtistDTOResponse{
			Id:          artist.Id,
			Name:        artist.Name,
			AvatarUrl:   artist.AvatarUrl,
			Description: artist.Description,
		})
	}

	return TrackDTOResponse{
		Id:              track.Id,
		Name:            track.Name,
		Description:     track.Description,
		AlbumId:         track.AlbumId,
		ArtistIds:       artistIds,
		Artists:         artists,
		GenreIds:        track.GenreIds,
		Genres:          track.Genres,
		CoverStorageKey: track.CoverStorageKey,
		DurationSecond:  int(track.DurationSecond.Seconds()),
		ReleaseDate:     track.ReleaseDate,
		IsExplicit:      track.IsExplicit,
		IsStreamable:    track.IsStreamable,
		IsDownloadable:  track.IsDownloadable,
		StatusId:        track.StatusId,
		Status:          track.Status,
		AudioStorageKey: track.AudioStorageKey,
	}
}
