package token_service

import (
	"context"
	"fmt"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
)

func (th *TokenService) SaveToken(ctx context.Context, token *domain.Token) (*domain.Token, error) {
	token, err := th.TokenRepository.SaveToken(ctx, token)

	if err != nil {
		return nil, fmt.Errorf("Failed removing token: %w", err)
	}

	return token, nil
}
