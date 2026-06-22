package artist_postgres_repository

import (
	"time"

	"github.com/google/uuid"
)

type ArtistModel struct {
	Id               uuid.UUID `json:"id"`
	Name             string    `json:"name"`
	Description      string    `json:"description"`
	AvatarStorageKey string    `json:"avatar_url"`
}

type AlbumModel struct {
	Id              uuid.UUID
	Title           string
	Description     string
	CoverStorageKey string
	ReleaseDate     time.Time
}
