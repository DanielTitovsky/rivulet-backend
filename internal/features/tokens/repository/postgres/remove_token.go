package tokens_postgres_repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (r *TokenRepository) RemoveToken(ctx context.Context, tokenId uuid.UUID) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.GetTimeout())
	defer cancel()

	query := `
       	DELETE 
		FROM refresh_tokens
		WHERE id = $1
		RETURNING id, user_id, token_hash, expires_at, revoked_at, created_at
    `

	_, err := r.pool.Exec(
		ctx,
		query,
		tokenId,
	)

	if err != nil {
		return fmt.Errorf("Failed delete token: %w", err)
	}

	return nil
}
