package tracks_postgres_repository

import (
	"context"
	"fmt"
	"time"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
	app_postgres_executor "github.com/DanielTitovsky/rivulet-backend.git/internal/app/repository/postgres"
	"github.com/google/uuid"
)

func (r *TrackRepository) UpdateTrack(ctx context.Context, trackId uuid.UUID, track domain.Track) error {
	execute := app_postgres_executor.GetQueryExecutor(ctx, r.pool)

	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)

	defer cancel()

	query := `
	UPDATE tracks
	SET
    title = $2,
    description = $3,
    album_id = $4,
    cover_url = $5,
    duration_seconds = $6,
    release_date = $7,
    is_explicit = $8,
    is_streamable = $9,
    is_downloadable = $10,
    status_id = $11,
    audio_storage_key = $12,
    updated_at = now()
	WHERE id = $1;
	`

	tag, err := execute.Exec(
		ctx,
		query,
		trackId,
		track.Name,
		track.Description,
		track.AlbumId,
		track.CoverStorageKey,
		int(track.DurationSecond.Seconds()),
		track.ReleaseDate,
		track.IsExplicit,
		track.IsStreamable,
		track.IsDownloadable,
		track.StatusId,
		track.AudioStorageKey,
	)

	if err != nil {
		return fmt.Errorf("update track: %w", err)
	}

	if tag.RowsAffected() == 0 {
		return fmt.Errorf("track not found or not updated")
	}

	return nil
}
