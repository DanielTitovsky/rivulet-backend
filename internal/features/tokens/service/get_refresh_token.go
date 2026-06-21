package tokens_service

import (
	"context"
	"fmt"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
)

func (s *TokensServise) GetRefreshToken(
	ctx context.Context,
	rawToken string,
) (domain.Token, error) {
	tokenHash := s.hashToken(rawToken)

	token, err := s.TokenRepository.GetTokenByHash(ctx, tokenHash)

	if err != nil {
		return domain.Token{}, fmt.Errorf("get refresh token: %w", err)
	}

	return token, nil
}
