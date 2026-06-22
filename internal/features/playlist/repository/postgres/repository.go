package playlist_postgres_repository

import app_postgres_pool "github.com/DanielTitovsky/rivulet-backend.git/internal/app/repository/postgres/pool"

type PlaylistRepository struct {
	pool app_postgres_pool.Pool
}

func NewUsersRepository(pool app_postgres_pool.Pool) *PlaylistRepository {
	return &PlaylistRepository{
		pool: pool,
	}
}
