package users_postgres_repository

import app_postgres_pool "github.com/DanielTitovsky/rivulet-backend.git/internal/app/repository/postgres/pool"

type UsersRepository struct {
	pool app_postgres_pool.Pool
}

func NewUsersRepository(pool app_postgres_pool.Pool) *UsersRepository {
	return &UsersRepository{
		pool: pool,
	}
}

// SaveUser(ctx context.Context, user domain.User) (domain.User, error)
// FindUserById(ctx context.Context, userId uuid.UUID) (*domain.User, error)
// UpdateUser(ctx context.Context, userId domain.User) (*domain.User, error)
// DeleteUser(ctx context.Context, userId uuid.UUID) error
