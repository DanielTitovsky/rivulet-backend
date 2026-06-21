package users_postgres_repository

import (
	"context"
	"fmt"
	"time"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
	app_postgres_executor "github.com/DanielTitovsky/rivulet-backend.git/internal/app/repository/postgres"
)

func (r *UsersRepository) CreateOAuthUser(
	ctx context.Context,
	email string,
	name string,
) (domain.User, error) {
	var userModel UserModel

	executor := app_postgres_executor.GetQueryExecutor(ctx, r.pool)

	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	query := `
		INSERT INTO users (
			email,
			name,
			hash_password
		)
		VALUES (
			$1,
			$2,
			NULL
		)
		RETURNING
			id,
			name,
			email,
			hash_password,
			created_at,
			updated_at
	`

	err := executor.QueryRow(
		ctx,
		query,
		email,
		name,
	).Scan(
		&userModel.Id,
		&userModel.Name,
		&userModel.Email,
		&userModel.Password,
		&userModel.CreatedAt,
		&userModel.UpdatedAt,
	)

	if err != nil {
		return domain.User{}, fmt.Errorf("create oauth user: %w", err)
	}

	return domain.User{
		Id:        userModel.Id,
		Name:      userModel.Name,
		Email:     userModel.Email,
		Password:  userModel.Password,
		CreatedAt: userModel.CreatedAt,
		UpdatedAt: userModel.UpdatedAt,
	}, nil
}
