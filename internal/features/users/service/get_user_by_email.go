package users_service

import (
	"context"
	"fmt"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
)

func (s *UsersServise) GetUserByEmail(ctx context.Context, email string) (domain.User, error) {
	user, err := s.UsersRepository.GetUserByEmail(ctx, email)

	if err != nil {
		return domain.User{}, fmt.Errorf("get user by email: %w", err)
	}

	return user, nil
}
