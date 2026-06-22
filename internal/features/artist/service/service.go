package artist_service

import (
	"context"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
	app_postgres_transaction "github.com/DanielTitovsky/rivulet-backend.git/internal/app/repository/postgres/transaction"
	"github.com/google/uuid"
)

type ArtistService struct {
	ArtistRepository   ArtistRepository
	ArtistStorage      ArtistStorage
	TransactionManager app_postgres_transaction.TransactionManager
}

type ArtistRepository interface {
	GetArtist(ctx context.Context, artistId uuid.UUID) (domain.Artist, error)
	GetArtistAlbums(ctx context.Context, artistId uuid.UUID) ([]domain.Album, error)
}
type ArtistStorage interface {
	GetArtistAvatar(ctx context.Context, link string) (string, error)
}

func NewArtistService(artistRepository ArtistRepository, artistStorage ArtistStorage, transactionManager app_postgres_transaction.TransactionManager) *ArtistService {
	return &ArtistService{
		ArtistRepository:   artistRepository,
		ArtistStorage:      artistStorage,
		TransactionManager: transactionManager,
	}
}
