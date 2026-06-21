package tokens_service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (s *TokensServise) RemoveToken(
	ctx context.Context,
	tokenId uuid.UUID,
) error {
	err := s.TokenRepository.RemoveToken(ctx, tokenId)

	if err != nil {
		return fmt.Errorf("remove token: %w", err)
	}

	return nil
}
