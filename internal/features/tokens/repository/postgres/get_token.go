package tokens_postgres_repository

import (
	"context"
	"fmt"
	"time"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
	app_postgres_executor "github.com/DanielTitovsky/rivulet-backend.git/internal/app/repository/postgres"
)

func (r *TokenRepository) GetTokenByHash(
	ctx context.Context,
	tokenHash string,
) (domain.Token, error) {
	var tokenModel TokenModel

	executor := app_postgres_executor.GetQueryExecutor(ctx, r.pool)

	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	query := `
		SELECT 
			id,
			user_id,
			token_hash,
			expires_at,
			revoked_at,
			created_at
		FROM refresh_tokens
		WHERE token_hash = $1
		  AND revoked_at IS NULL
	`

	err := executor.QueryRow(
		ctx,
		query,
		tokenHash,
	).Scan(
		&tokenModel.Id,
		&tokenModel.UserId,
		&tokenModel.TokenString,
		&tokenModel.ExpiresAt,
		&tokenModel.RevokedAt,
		&tokenModel.CreatedAt,
	)

	if err != nil {
		return domain.Token{}, fmt.Errorf("get token by hash: %w", err)
	}

	return domain.Token{
		Id:          tokenModel.Id,
		UserId:      tokenModel.UserId,
		TokenString: tokenModel.TokenString,
		ExpiresAt:   tokenModel.ExpiresAt,
		RevokedAt:   tokenModel.RevokedAt,
		CreatedAt:   tokenModel.CreatedAt,
	}, nil
}
