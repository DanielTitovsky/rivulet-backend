package tokens_postgres_repository

import app_postgres_pool "github.com/DanielTitovsky/rivulet-backend.git/internal/app/repository/postgres/pool"

type TokenRepository struct {
	pool app_postgres_pool.Pool
}

func NewUsersRepository(pool app_postgres_pool.Pool) *TokenRepository {
	return &TokenRepository{
		pool: pool,
	}
}
