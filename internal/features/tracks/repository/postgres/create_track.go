package tracks_postgres_repository

import (
	"context"
	"fmt"
	"time"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
	app_postgres_executor "github.com/DanielTitovsky/rivulet-backend.git/internal/app/repository/postgres"
	"github.com/google/uuid"
)

func (r *TrackRepository) CreateTrack(ctx context.Context, track domain.Track) (uuid.UUID, error) {
	var trackId uuid.UUID
	executor := app_postgres_executor.GetQueryExecutor(ctx, r.pool)

	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	query := `
	INSERT INTO tracks (
	    id,
	    album_id,
	    title,
	    description,
	    cover_url,
	    duration_seconds,
	    release_date,
	    is_explicit,
	    is_streamable,
	    is_downloadable,
	    audio_storage_key,
	    status_id
	)
	VALUES (
	    $1,
	    $2,
	    $3,
	    $4,
	    $5,
	    $6,
	    $7,
	    $8,
	    $9,
	    $10,
	    $11,
	    $12
	)
	RETURNING  id;
	`

	err := executor.QueryRow(
		ctx,
		query,
		track.Id,
		track.AlbumId,
		track.Name,
		track.Description,
		track.CoverStorageKey,
		int(track.DurationSecond.Seconds()),
		track.ReleaseDate,
		track.IsExplicit,
		track.IsStreamable,
		track.IsDownloadable,
		track.AudioStorageKey,
		track.StatusId,
	).Scan(&trackId)

	if err != nil {
		return uuid.Nil, fmt.Errorf("Create track: %w", err)
	}

	return trackId, nil
}
