package artist_minio_repository

import app_minIo_storage "github.com/DanielTitovsky/rivulet-backend.git/internal/app/repository/minIo/storage"

type ArtistStorageRepository struct {
	storage app_minIo_storage.MinioStorage
}

func NewArtistStorageRepository(storage app_minIo_storage.MinioStorage) *ArtistStorageRepository {
	return &ArtistStorageRepository{
		storage: storage,
	}
}
