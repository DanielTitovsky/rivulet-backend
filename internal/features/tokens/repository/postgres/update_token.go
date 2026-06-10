package tokens_postgres_repository

import (
	"context"
	"fmt"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
)

func (r *TokenRepository) UpdateUserToken(ctx context.Context, token domain.Token) (*domain.Token, error) {
	var tokenModel TokenModel
	ctx, cancel := context.WithTimeout(ctx, r.pool.GetTimeout())
	defer cancel()

	query := `
		UPDATE refresh_tokens
		SET  token_hash = $2, expires_at = $3, revoked_at = $4
		WHERE id = $1
		RETURNING id, user_id, token_hash, expires_at, revoked_at, created_at
	`

	err := r.pool.QueryRow(
		ctx,
		query,
		token.Id,
		token.TokenString,
		token.ExpiresAt,
		token.RevokedAt,
	).Scan(
		&tokenModel.Id,
		&tokenModel.UserId,
		&tokenModel.TokenString,
		&tokenModel.ExpiresAt,
		&tokenModel.RevokedAt,
		&tokenModel.CreatedAt,
	)

	if err != nil {
		return &domain.Token{}, fmt.Errorf("Failed save token: %w", err)
	}

	return &domain.Token{
		Id:          tokenModel.Id,
		UserId:      tokenModel.UserId,
		TokenString: tokenModel.TokenString,
		ExpiresAt:   tokenModel.ExpiresAt,
		RevokedAt:   tokenModel.RevokedAt,
		CreatedAt:   tokenModel.CreatedAt,
	}, nil
}
