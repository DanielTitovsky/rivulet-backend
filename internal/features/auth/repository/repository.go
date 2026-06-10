package auth_postgres_repository

import app_postgres_pool "github.com/DanielTitovsky/rivulet-backend.git/internal/app/repository/postgres/pool"

type AuthRepository struct {
	pool app_postgres_pool.Pool
}

func NewAuthRepository(pool app_postgres_pool.Pool) *AuthRepository {
	return &AuthRepository{
		pool: pool,
	}
}
