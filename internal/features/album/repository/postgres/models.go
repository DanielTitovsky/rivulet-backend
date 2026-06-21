package artist_postgres_repository

import (
	"github.com/google/uuid"
)

type AlbomModel struct {
	Id               uuid.UUID   `json:"id"`
	Name             string      `json:"name"`
	ArtistIds        []uuid.UUID `json:"artistIds"`
	Description      string      `json:"description"`
	AvatarStorageKey string      `json:"avatar_url"`
}

type ArtistModel struct {
	Id       uuid.UUID
	Name     string
	CoverUrl string
}
