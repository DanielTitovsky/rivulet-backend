package tracks_postgres_repository

import (
	"context"
	"fmt"
	"time"

	app_postgres_executor "github.com/DanielTitovsky/rivulet-backend.git/internal/app/repository/postgres"
	"github.com/google/uuid"
)

func (r *TrackRepository) UpdateTrackGenres(ctx context.Context, trackId uuid.UUID, genreId []uuid.UUID) error {
	executor := app_postgres_executor.GetQueryExecutor(ctx, r.pool)

	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)

	defer cancel()

	deleteQuery := `
		DELETE FROM track_genres
		WHERE track_id = $1;
	`

	_, err := executor.Exec(ctx, deleteQuery, trackId)
	if err != nil {
		return fmt.Errorf("delete track_genres: %w", err)
	}

	insertQuery := `
		INSERT INTO track_genres (track_id, genre_id)
		SELECT
    	$1,
    	unnest($2::uuid[]);
	`

	_, err = executor.Exec(ctx, insertQuery, trackId, genreId)
	if err != nil {
		return fmt.Errorf("insert track_genres: %w", err)
	}

	return nil
}
