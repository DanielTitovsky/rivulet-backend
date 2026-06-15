package artist_postgres_repository

import (
	"github.com/google/uuid"
)

type ArtistModel struct {
	Id               uuid.UUID `json:"id"`
	Name             string    `json:"name"`
	Description      string    `json:"description"`
	AvatarStorageKey string    `json:"avatar_url"`
}
