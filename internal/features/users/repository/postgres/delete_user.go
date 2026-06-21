package users_postgres_repository

import (
	"context"
	"fmt"
	"time"

	app_postgres_executor "github.com/DanielTitovsky/rivulet-backend.git/internal/app/repository/postgres"
	"github.com/google/uuid"
)

func (r *UsersRepository) DeleteUser(ctx context.Context, userId uuid.UUID) error {
	executor := app_postgres_executor.GetQueryExecutor(ctx, r.pool)

	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	query := `
		DELETE 
		FROM users
		WHERE id = $1
	`
	_, err := executor.Exec(ctx, query, userId)

	if err != nil {
		return fmt.Errorf("Delete user: %w", err)
	}

	return nil
}
