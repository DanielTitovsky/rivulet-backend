package users_postgres_repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (r *UsersRepository) DeleteUser(ctx context.Context, userId uuid.UUID) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.GetTimeout())

	defer cancel()

	query := `
		DELETE 
		FROM users
		WHERE id = $1
	`
	_, err := r.pool.Exec(ctx, query, userId)

	if err != nil {
		return fmt.Errorf("Delete user: %w", err)
	}

	return nil
}
