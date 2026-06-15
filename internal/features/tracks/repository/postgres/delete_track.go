package tracks_postgres_repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (r *TrackRepository) DeleteTrack(ctx context.Context, trackId uuid.UUID) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.GetTimeout())
	defer cancel()

	query := `
	DELETE FROM tracks 
	WHERE tracks.id = $1
	`

	_, err := r.pool.Exec(ctx, query, trackId)

	if err != nil {
		return fmt.Errorf("Delete track: %w", err)
	}

	return nil
}
