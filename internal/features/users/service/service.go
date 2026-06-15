package users_service

import (
	"context"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
	app_postgres_transaction "github.com/DanielTitovsky/rivulet-backend.git/internal/app/repository/postgres/transaction"
	"github.com/google/uuid"
)

type UsersServise struct {
	UsersRepository    UsersRepository
	TransactionManager app_postgres_transaction.TransactionManager
}

type UsersRepository interface {
	SaveUser(ctx context.Context, user domain.User) (domain.User, error)
	GetUser(ctx context.Context, userId uuid.UUID) (domain.User, error)
	UpdateUser(ctx context.Context, userId uuid.UUID, user domain.User) (domain.User, error)
	DeleteUser(ctx context.Context, userId uuid.UUID) error
}

func NewUserServise(rep UsersRepository, transactionManager app_postgres_transaction.TransactionManager) *UsersServise {
	return &UsersServise{
		UsersRepository:    rep,
		TransactionManager: transactionManager,
	}
}
