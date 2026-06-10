package token_service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (th *TokenService) RemoveToken(ctx context.Context, tokenId uuid.UUID) error {
	err := th.TokenRepository.RemoveToken(ctx, tokenId)

	if err != nil {
		return fmt.Errorf("Failed removing token: %w", err)
	}

	return nil
}
