package tracks_postgres_repository

import (
	"context"
	"fmt"
	"time"

	app_postgres_executor "github.com/DanielTitovsky/rivulet-backend.git/internal/app/repository/postgres"
	"github.com/google/uuid"
)

func (r *TrackRepository) CreateTrackGenres(ctx context.Context, trackId uuid.UUID, genreIds []uuid.UUID) error {
	executor := app_postgres_executor.GetQueryExecutor(ctx, r.pool)

	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	query := `
	INSERT INTO track_genres (
    	track_id,
    	genre_id
	)
	SELECT
    	$1,
    	unnest($2::uuid[]);
	`

	_, err := executor.Exec(ctx, query, trackId, genreIds)

	if err != nil {
		return fmt.Errorf("Create track_genres: %w", err)
	}

	return nil
}
