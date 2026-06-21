package artist_postgres_repository

import (
	"github.com/google/uuid"
)

type PlaylistModel struct {
	Id               uuid.UUID `json:"id"`
	User             UserModel
	Name             string `json:"name"`
	Description      string `json:"description"`
	AvatarStorageKey string `json:"avatar_url"`
}

type UserModel struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	AvatarUrl string    `json:"avatar_url"`
}
