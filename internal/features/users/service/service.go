package users_service

import (
	"context"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
)

type UsersServise struct {
	UsersRepository
}

type UsersRepository interface {
	SaveUser(ctx context.Context, user domain.User) (domain.User, error)
}

// FindUserById(ctx context.Context, userId uuid.UUID) (*domain.User, error)
// 	FindUserByEmail(ctx context.Context, userEmail string) (*domain.User, error)
// 	SaveOAthAccaounts(ctx context.Context, provideUser domain.ProvideUser, userId uuid.UUID) error
// 	UpdateUser(ctx context.Context, userId domain.User) (*domain.User, error)
// 	DeleteUser(ctx context.Context, userId uuid.UUID) error

func NewUserServise(rep UsersRepository) *UsersServise {
	return &UsersServise{
		UsersRepository: rep,
	}
}
