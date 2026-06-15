package tracks_postgres_repository

import (
	"context"
	"fmt"
	"time"

	app_postgres_executor "github.com/DanielTitovsky/rivulet-backend.git/internal/app/repository/postgres"
	"github.com/google/uuid"
)

func (r *TrackRepository) UpdateTrackArtists(ctx context.Context, trackId uuid.UUID, artistsId []uuid.UUID) error {
	executor := app_postgres_executor.GetQueryExecutor(ctx, r.pool)

	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)

	defer cancel()

	deleteQuery := `
		DELETE FROM track_artists
		WHERE track_id = $1;
	`

	_, err := executor.Exec(ctx, deleteQuery, trackId)
	if err != nil {
		return fmt.Errorf("delete track_artists: %w", err)
	}

	insertQuery := `
	INSERT INTO track_artists (track_id, artist_id, role)
	SELECT
    $1,
    unnest($2::uuid[]),
    'artist';
	`

	_, err = executor.Exec(ctx, insertQuery, trackId, artistsId)
	if err != nil {
		return fmt.Errorf("insert track_artists: %w", err)
	}

	return nil
}
