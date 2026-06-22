package playlist_postgres_repository

import (
	"github.com/google/uuid"
)

type PlaylistModel struct {
	Id          uuid.UUID `json:"id"`
	UserId      *uuid.UUID
	Name        string `json:"name"`
	CoverUrl    string `json:"cover_url"`
	Description string `json:"description"`
	Visibility  bool   `json:"visibility"`
}

type UserModel struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	AvatarUrl string    `json:"avatar_url"`
}
