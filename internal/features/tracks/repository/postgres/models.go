package tracks_postgres_repository

import (
	"time"

	"github.com/google/uuid"
)

type TrackModel struct {
	Id              uuid.UUID
	AlbumId         *uuid.UUID
	Name            string
	Description     string
	CoverStorageKey string
	DurationSecond  int
	ReleaseDate     time.Time
	IsExplicit      bool
	IsStreamable    bool
	IsDownloadable  bool
	AudioStorageKey string
	Status          string
	Genres          []string
	ArtistsJSON     []byte
	GenresJSON      []byte
}

type trackArtistModel struct {
	Id               uuid.UUID `json:"id"`
	Name             string    `json:"name"`
	Description      string    `json:"description"`
	AvatarStorageKey string    `json:"avatar_url"`
}
