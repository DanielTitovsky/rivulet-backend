package users_postgres_repository

import (
	"context"
	"fmt"
	"time"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
	"github.com/google/uuid"
)

func (r *UsersRepository) GetUserById(ctx context.Context, userId uuid.UUID) (domain.User, error) {
	var userModel UserModel
	ctx, cancel := context.WithTimeout(ctx, r.pool.GetTimeout())
	defer cancel()

	query := `
	    SELECT id, name, email
		FROM users
		WHERE id = $1
	`

	err := r.pool.QueryRow(
		ctx,
		query,
		userId,
	).Scan(
		&userModel.Id,
		&userModel.Name,
		&userModel.Email,
	)

	if err != nil {
		fmt.Print("\n")
		fmt.Print("\n")
		fmt.Print(err)
		fmt.Print("\n")
		fmt.Print("\n")
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
