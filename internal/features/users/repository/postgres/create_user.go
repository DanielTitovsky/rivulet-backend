package users_postgres_repository

import (
	"context"
	"fmt"
	"time"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
	app_postgres_executor "github.com/DanielTitovsky/rivulet-backend.git/internal/app/repository/postgres"
)

func (r *UsersRepository) SaveUser(ctx context.Context, user domain.User) (domain.User, error) {
	var userModel UserModel
	executor := app_postgres_executor.GetQueryExecutor(ctx, r.pool)

	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	query := `
        INSERT INTO users (
            email, 
            name, 
            hash_password, 
            updated_at, 
            created_at
        ) 
        VALUES ($1, $2, $3, $4, $5)
		RETURNING id, name, email
    `

	err := executor.QueryRow(
		ctx,
		query,
		user.Email,
		user.Name,
		*user.Password,
		user.UpdatedAt,
		user.CreatedAt,
	).Scan(
		&userModel.Id,
		&userModel.Name,
		&userModel.Email,
	)

	if err != nil {
		return domain.User{}, fmt.Errorf("Scan user: %w", err)
	}

	return domain.User{
		Id:        userModel.Id,
		Name:      userModel.Name,
		Email:     userModel.Email,
		Password:  nil,
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}, nil
}
