package tokens_postgres_repository

import (
	"context"
	"fmt"
	"time"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
	app_postgres_executor "github.com/DanielTitovsky/rivulet-backend.git/internal/app/repository/postgres"
)

func (r *TokenRepository) SaveToken(
	ctx context.Context,
	token domain.Token,
) (domain.Token, error) {
	var tokenModel TokenModel

	executor := app_postgres_executor.GetQueryExecutor(ctx, r.pool)

	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	query := `
		INSERT INTO refresh_tokens (
			id,
			user_id,
			token_hash,
			expires_at,
			revoked_at,
			created_at
		)
		VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT (user_id)
		DO UPDATE SET
			id = EXCLUDED.id,
			token_hash = EXCLUDED.token_hash,
			expires_at = EXCLUDED.expires_at,
			revoked_at = EXCLUDED.revoked_at,
			created_at = EXCLUDED.created_at
		RETURNING id, user_id, token_hash, expires_at, revoked_at, created_at
	`

	err := executor.QueryRow(
		ctx,
		query,
		token.Id,
		token.UserId,
		token.TokenString,
		token.ExpiresAt,
		token.RevokedAt,
		token.CreatedAt,
	).Scan(
		&tokenModel.Id,
		&tokenModel.UserId,
		&tokenModel.TokenString,
		&tokenModel.ExpiresAt,
		&tokenModel.RevokedAt,
		&tokenModel.CreatedAt,
	)

	if err != nil {
		return domain.Token{}, fmt.Errorf("save token: %w", err)
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
