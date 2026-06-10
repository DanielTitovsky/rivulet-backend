package users_service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (s *UsersServise) DeleteUser(ctx context.Context, UserId uuid.UUID) error {

	err := s.UsersRepository.DeleteUser(ctx, UserId)

	if err != nil {
		return fmt.Errorf("Create user: %w", err)
	}

	return nil
}
