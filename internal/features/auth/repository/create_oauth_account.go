package auth_postgres_repository

import (
	"context"
	"fmt"
	"time"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
	app_postgres_executor "github.com/DanielTitovsky/rivulet-backend.git/internal/app/repository/postgres"
	"github.com/google/uuid"
)

func (r *AuthRepository) CreateOAuthAccount(
	ctx context.Context,
	userId uuid.UUID,
	oauthUser domain.OAuthUser,
) error {
	executor := app_postgres_executor.GetQueryExecutor(ctx, r.pool)

	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	query := `
		INSERT INTO oauth_accounts (
			user_id,
			provider,
			provider_user_id
		)
		VALUES (
			$1,
			$2,
			$3
		)
	`

	_, err := executor.Exec(
		ctx,
		query,
		userId,
		string(oauthUser.Provider),
		oauthUser.ProviderUserId,
	)

	if err != nil {
		return fmt.Errorf("create oauth account: %w", err)
	}

	return nil
}
