package users_service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (s *UsersServise) RemoveTrackFromFavorite(
	ctx context.Context,
	userId uuid.UUID,
	trackId uuid.UUID,
) error {

	err := s.UsersRepository.RemoveTrackFromFavorite(ctx, userId, trackId)

	if err != nil {
		return fmt.Errorf("remove track from favorite: %w", err)
	}

	return nil
}
