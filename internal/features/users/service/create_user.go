package users_service

import (
	"context"
	"fmt"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
)

func (s *UsersServise) CreateUser(ctx context.Context, user domain.User) (domain.User, error) {

	if err := user.Validate(); err != nil {
		return domain.User{}, fmt.Errorf("Validate user domain: %w", err)
	}

	hashedPassword := user.HashUserPassword(*user.Password)

	user.Password = &hashedPassword

	user, err := s.UsersRepository.SaveUser(ctx, user)

	if err != nil {
		return domain.User{}, fmt.Errorf("Create user: %w", err)
	}

	return user, nil
}
