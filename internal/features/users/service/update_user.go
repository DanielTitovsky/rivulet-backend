package users_service

import (
	"context"
	"fmt"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
	"github.com/google/uuid"
)

func (s *UsersServise) UpdateUser(ctx context.Context, userId uuid.UUID, updateUser domain.UserUpdate) (domain.User, error) {

	if err := updateUser.Validate(); err != nil {
		return domain.User{}, fmt.Errorf("Validate user domain: %w", err)
	}

	user, err := s.UsersRepository.GetUserById(ctx, userId)

	if err != nil {
		return domain.User{}, fmt.Errorf("get user: %w", err)
	}

	if err := user.ApplyUpdate(updateUser); err != nil {
		return domain.User{}, fmt.Errorf("apply user update: %w", err)
	}

	hashedPassword := user.HashUserPassword(updateUser.Password)

	user.Password = &hashedPassword

	updatedUser, err := s.UsersRepository.UpdateUser(ctx, userId, user)

	if err != nil {
		return domain.User{}, fmt.Errorf("Create user: %w", err)
	}

	return updatedUser, nil
}
