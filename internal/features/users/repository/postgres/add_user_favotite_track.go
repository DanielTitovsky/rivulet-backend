package users_postgres_repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (r *UsersRepository) AddTrackToFavorite(
	ctx context.Context,
	userId uuid.UUID,
	trackId uuid.UUID,
) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.GetTimeout())
	defer cancel()

	query := `
		INSERT INTO user_liked_tracks (
			user_id,
			track_id
		)
		VALUES ($1, $2)
		ON CONFLICT (user_id, track_id) DO NOTHING
	`

	_, err := r.pool.Exec(
		ctx,
		query,
		userId,
		trackId,
	)

	if err != nil {
		return fmt.Errorf("insert user liked track: %w", err)
	}

	return nil
}
