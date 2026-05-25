package users_postgres_repository

import app_postgres_pool "github.com/DanielTitovsky/rivulet-backend.git/internal/app/repository/postgres/pool"

type UsersRepository struct {
	pool app_postgres_pool.Pool
}

//7:22:33

func NewUsersRepository(pool app_postgres_pool.Pool) *UsersRepository {
	return &UsersRepository{
		pool: pool,
	}
}
