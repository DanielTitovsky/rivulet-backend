package tracks_minio_repository

import (
	app_minIo_storage "github.com/DanielTitovsky/rivulet-backend.git/internal/app/repository/minIo/storage"
)

type TrackStorageRepository struct {
	storage app_minIo_storage.Storage
}

func NewTrackRepository(storage app_minIo_storage.Storage) *TrackStorageRepository {
	return &TrackStorageRepository{
		storage: storage,
	}
}
