package users_postgres_repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (r *UsersRepository) RemoveTrackFromFavorite(
	ctx context.Context,
	userId uuid.UUID,
	trackId uuid.UUID,
) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.GetTimeout())
	defer cancel()

	query := `
		DELETE FROM user_liked_tracks
		WHERE user_id = $1
		  AND track_id = $2
	`

	_, err := r.pool.Exec(
		ctx,
		query,
		userId,
		trackId,
	)

	if err != nil {
		return fmt.Errorf("delete user liked track: %w", err)
	}

	return nil
}
