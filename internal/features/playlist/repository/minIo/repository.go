package playlis_minio_repository

import (
	app_minIo_storage "github.com/DanielTitovsky/rivulet-backend.git/internal/app/repository/minIo/storage"
)

type PlaylistStorageRepository struct {
	storage app_minIo_storage.Storage
}

func NewPlaylistRepository(storage app_minIo_storage.Storage) *PlaylistStorageRepository {
	return &PlaylistStorageRepository{
		storage: storage,
	}
}
