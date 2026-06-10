package token_service

import (
	"context"
	"fmt"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
	"github.com/google/uuid"
)

func (th *TokenService) GetToken(ctx context.Context, id uuid.UUID) (*domain.Token, error) {
	token, err := th.TokenRepository.GetToken(ctx, id)

	if err != nil {
		return nil, fmt.Errorf("Failed getting token: %w", err)
	}

	return token, nil
}
