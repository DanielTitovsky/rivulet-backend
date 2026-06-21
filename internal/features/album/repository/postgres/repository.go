package artist_postgres_repository

import app_postgres_pool "github.com/DanielTitovsky/rivulet-backend.git/internal/app/repository/postgres/pool"

type ArtistRepository struct {
	pool app_postgres_pool.Pool
}

func NewUsersRepository(pool app_postgres_pool.Pool) *ArtistRepository {
	return &ArtistRepository{
		pool: pool,
	}
}
