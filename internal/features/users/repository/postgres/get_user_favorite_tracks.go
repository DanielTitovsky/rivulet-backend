package users_postgres_repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
	app_postgres_executor "github.com/DanielTitovsky/rivulet-backend.git/internal/app/repository/postgres"
	"github.com/google/uuid"
)

func (r *UsersRepository) GetUserFavoriteTracks(
	ctx context.Context,
	userId uuid.UUID,
) ([]domain.Track, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.GetTimeout())
	defer cancel()

	executor := app_postgres_executor.GetQueryExecutor(ctx, r.pool)

	query := `
SELECT
	t.id,
	t.album_id,
	t.title,
	COALESCE(t.description, ''),
	COALESCE(t.cover_url, ''),
	t.duration_seconds,
	COALESCE(t.release_date, '1970-01-01'::date),
	t.is_explicit,
	t.is_streamable,
	t.is_downloadable,
	COALESCE(t.audio_storage_key, ''),
	ts.code,
	COALESCE(
		json_agg(
			DISTINCT jsonb_build_object(
				'id', a.id,
				'name', a.name,
				'description', COALESCE(a.description, ''),
				'avatar_url', COALESCE(a.avatar_url, '')
			)
		) FILTER (WHERE a.id IS NOT NULL),
		'[]'::json
	) AS artists_json,
	COALESCE(
		json_agg(DISTINCT g.name) FILTER (WHERE g.name IS NOT NULL),
		'[]'::json
	) AS genres_json
FROM user_liked_tracks ult
JOIN tracks t
	ON t.id = ult.track_id
LEFT JOIN track_statuses ts
	ON ts.id = t.status_id
LEFT JOIN track_artists ta
	ON ta.track_id = t.id
LEFT JOIN artists a
	ON a.id = ta.artist_id
LEFT JOIN track_genres tg
	ON tg.track_id = t.id
LEFT JOIN genres g
	ON g.id = tg.genre_id
WHERE ult.user_id = $1
GROUP BY
	t.id,
	t.album_id,
	t.title,
	t.description,
	t.cover_url,
	t.duration_seconds,
	t.release_date,
	t.is_explicit,
	t.is_streamable,
	t.is_downloadable,
	t.audio_storage_key,
	ts.code,
	ult.created_at
ORDER BY ult.created_at DESC;
`

	rows, err := executor.Query(ctx, query, userId)

	if err != nil {
		return nil, fmt.Errorf("Select user favorite tracks: %w", err)
	}

	defer rows.Close()

	tracks := make([]domain.Track, 0)

	for rows.Next() {
		var trackModel TrackModel

		err = rows.Scan(
			&trackModel.Id,
			&trackModel.AlbumId,
			&trackModel.Name,
			&trackModel.Description,
			&trackModel.CoverStorageKey,
			&trackModel.DurationSecond,
			&trackModel.ReleaseDate,
			&trackModel.IsExplicit,
			&trackModel.IsStreamable,
			&trackModel.IsDownloadable,
			&trackModel.AudioStorageKey,
			&trackModel.Status,
			&trackModel.ArtistsJSON,
			&trackModel.GenresJSON,
		)

		if err != nil {
			return nil, fmt.Errorf("Scan user favorite track: %w", err)
		}

		var artists []domain.Artist
		if err := json.Unmarshal(trackModel.ArtistsJSON, &artists); err != nil {
			return nil, fmt.Errorf("Unmarshal favorite track artists: %w", err)
		}

		var genres []string
		if err := json.Unmarshal(trackModel.GenresJSON, &genres); err != nil {
			return nil, fmt.Errorf("Unmarshal favorite track genres: %w", err)
		}

		tracks = append(tracks, domain.Track{
			Id:              trackModel.Id,
			AlbumId:         trackModel.AlbumId,
			Name:            trackModel.Name,
			Description:     trackModel.Description,
			Artists:         artists,
			Genres:          genres,
			CoverStorageKey: trackModel.CoverStorageKey,
			DurationSecond:  time.Duration(trackModel.DurationSecond) * time.Second,
			ReleaseDate:     trackModel.ReleaseDate,
			IsExplicit:      trackModel.IsExplicit,
			IsStreamable:    trackModel.IsStreamable,
			IsDownloadable:  trackModel.IsDownloadable,
			Status:          trackModel.Status,
			AudioStorageKey: trackModel.AudioStorageKey,
		})
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("Iterate user favorite tracks: %w", err)
	}

	return tracks, nil
}
