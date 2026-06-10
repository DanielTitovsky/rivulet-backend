package users_postgres_repository

import (
	"context"
	"fmt"
	"time"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
	"github.com/google/uuid"
)

func (r *UsersRepository) UpdateUser(ctx context.Context, userId uuid.UUID, user domain.User) (domain.User, error) {
	var userModel UserModel

	ctx, cancel := context.WithTimeout(ctx, r.pool.GetTimeout())
	defer cancel()

	query := `
		UPDATE users
		SET email = $2, name = $3, hash_password = $4, updated_at = $5
		WHERE id = $1
		RETURNING id, name, email
	`

	err := r.pool.QueryRow(
		ctx,
		query,
		userId,
		user.Email,
		user.Name,
		user.Password,
		time.Now(),
	).Scan(
		&userModel.Id,
		&userModel.Name,
		&userModel.Email,
	)

	if err != nil {
		fmt.Print("\n")
		fmt.Print(err)
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
