package tokens_postgres_repository

import (
	"context"
	"fmt"
	"time"

	app_postgres_executor "github.com/DanielTitovsky/rivulet-backend.git/internal/app/repository/postgres"
	"github.com/google/uuid"
)

func (r *TokenRepository) RemoveToken(
	ctx context.Context,
	tokenId uuid.UUID,
) error {
	executor := app_postgres_executor.GetQueryExecutor(ctx, r.pool)

	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	query := `
		DELETE FROM refresh_tokens
		WHERE id = $1
	`

	_, err := executor.Exec(ctx, query, tokenId)

	if err != nil {
		return fmt.Errorf("remove token: %w", err)
	}

	return nil
}
