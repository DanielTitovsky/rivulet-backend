package users_service

import (
	"context"
	"fmt"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
	"github.com/google/uuid"
)

func (s *UsersServise) GetUser(ctx context.Context, userId uuid.UUID) (domain.User, error) {
	user, err := s.UsersRepository.GetUserById(ctx, userId)

	if err != nil {
		return domain.User{}, fmt.Errorf("Get user: %w", err)
	}

	return user, nil
}
