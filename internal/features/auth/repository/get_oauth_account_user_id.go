package auth_postgres_repository

import (
	"context"
	"fmt"
	"time"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
	app_postgres_executor "github.com/DanielTitovsky/rivulet-backend.git/internal/app/repository/postgres"
	"github.com/google/uuid"
)

func (r *AuthRepository) GetOAuthAccountUserId(
	ctx context.Context,
	provider domain.ProviderType,
	providerUserId string,
) (uuid.UUID, error) {
	var userId uuid.UUID

	executor := app_postgres_executor.GetQueryExecutor(ctx, r.pool)

	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	query := `
		SELECT
			user_id
		FROM oauth_accounts
		WHERE provider = $1
		  AND provider_user_id = $2
	`

	err := executor.QueryRow(
		ctx,
		query,
		string(provider),
		providerUserId,
	).Scan(
		&userId,
	)

	if err != nil {
		return uuid.Nil, fmt.Errorf("get oauth account user id: %w", err)
	}

	return userId, nil
}
