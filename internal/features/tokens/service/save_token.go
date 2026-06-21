package tokens_service

import (
	"context"
	"fmt"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
)

func (s *TokensServise) SaveRefreshToken(
	ctx context.Context,
	token domain.Token,
) (domain.Token, error) {
	token.TokenString = s.hashToken(token.TokenString)

	savedToken, err := s.TokenRepository.SaveToken(ctx, token)

	if err != nil {
		return domain.Token{}, fmt.Errorf("save refresh token: %w", err)
	}

	return savedToken, nil
}
