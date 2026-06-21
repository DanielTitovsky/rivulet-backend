package users_service

import (
	"context"
	"errors"
	"fmt"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
	"github.com/jackc/pgx/v5"
)

func (s *UsersServise) GetOrCreateOAuthUser(
	ctx context.Context,
	email string,
	name string,
) (domain.User, error) {
	user, err := s.UsersRepository.GetUserByEmail(ctx, email)

	if err == nil {
		return user, nil
	}

	if !errors.Is(err, pgx.ErrNoRows) {
		return domain.User{}, fmt.Errorf("get user by email: %w", err)
	}

	createdUser, err := s.UsersRepository.CreateOAuthUser(ctx, email, name)
	if err != nil {
		return domain.User{}, fmt.Errorf("create oauth user: %w", err)
	}

	return createdUser, nil
}
