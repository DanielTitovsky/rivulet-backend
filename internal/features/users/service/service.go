package users_service

import (
	"context"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
	"github.com/google/uuid"
)

type UsersServise struct {
	UsersRepository
}

type UsersRepository interface {
	SaveUser(ctx context.Context, user domain.User) (domain.User, error)
	GetUserById(ctx context.Context, userId uuid.UUID) (domain.User, error)
	UpdateUser(ctx context.Context, userId uuid.UUID, user domain.User) (domain.User, error)
	DeleteUser(ctx context.Context, userId uuid.UUID) error
}

// 	FindUserByEmail(ctx context.Context, userEmail string) (*domain.User, error)

func NewUserServise(rep UsersRepository) *UsersServise {
	return &UsersServise{
		UsersRepository: rep,
	}
}
