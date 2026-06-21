package users_service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (s *UsersServise) AddTrackToFavorite(
	ctx context.Context,
	userId uuid.UUID,
	trackId uuid.UUID,
) error {
	err := s.UsersRepository.AddTrackToFavorite(ctx, userId, trackId)

	if err != nil {
		return fmt.Errorf("add track to favorite: %w", err)
	}

	return nil
}
