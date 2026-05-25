package users_postgres_repository

import (
	"context"
	"fmt"
	"time"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
)

func (r *UsersRepository) SaveUser(ctx context.Context, user domain.User) (domain.User, error) {
	var userModel UserModel
	ctx, cancel := context.WithTimeout(ctx, r.pool.GetTimeout())
	defer cancel()

	query := `
        INSERT INTO users (
            id, 
            email, 
            name, 
            hash_password, 
            updated_at, 
            created_at
        ) 
        VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, name, email
    `

	err := r.pool.QueryRow(ctx, query, user.Id, user.Email, user.Name, *user.Password, user.UpdatedAt, user.CreatedAt).Scan(&userModel.Id, &userModel.Name, &userModel.Email)

	if err != nil {
		fmt.Print(err)
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
