package users_postgres_repository

import (
	"time"

	"github.com/google/uuid"
)

type UserModel struct {
	Id        uuid.UUID
	Email     *string
	Name      string
	Password  *string
	UpdatedAt time.Time
	CreatedAt time.Time
}
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
	ArtistsJSON     []byte
	GenresJSON      []byte
}
