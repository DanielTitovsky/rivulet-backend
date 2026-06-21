package users_postgres_repository

import (
	"context"
	"fmt"
	"time"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
	app_postgres_executor "github.com/DanielTitovsky/rivulet-backend.git/internal/app/repository/postgres"
)

func (r *UsersRepository) GetUserByEmail(ctx context.Context, email string) (domain.User, error) {
	var userModel UserModel
	executor := app_postgres_executor.GetQueryExecutor(ctx, r.pool)

	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	query := `
	    SELECT id, name, email, hash_password, created_at, updated_at
		FROM users
		WHERE email = $1
	`

	err := executor.QueryRow(
		ctx,
		query,
		email,
	).Scan(
		&userModel.Id,
		&userModel.Name,
		&userModel.Email,
		&userModel.Password,
		&userModel.CreatedAt,
		&userModel.UpdatedAt,
	)

	if err != nil {
		return domain.User{}, fmt.Errorf("scan user by email: %w", err)
	}

	return domain.User{
		Id:        userModel.Id,
		Name:      userModel.Name,
		Email:     userModel.Email,
		Password:  userModel.Password,
		UpdatedAt: userModel.UpdatedAt,
		CreatedAt: userModel.CreatedAt,
	}, nil
}
