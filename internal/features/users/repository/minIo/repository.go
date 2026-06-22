package users_minio_repository

import (
	app_minIo_storage "github.com/DanielTitovsky/rivulet-backend.git/internal/app/repository/minIo/storage"
)

type UserStorageRepository struct {
	storage app_minIo_storage.Storage
}

func NewUserRepository(storage app_minIo_storage.Storage) *UserStorageRepository {
	return &UserStorageRepository{
		storage: storage,
	}
}
