package tracks_postgres_repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
	app_postgres_executor "github.com/DanielTitovsky/rivulet-backend.git/internal/app/repository/postgres"
	"github.com/google/uuid"
)

func (r *TrackRepository) GetTracksByArtistId(
	ctx context.Context,
	artistId uuid.UUID,
) ([]domain.Track, error) {
	executor := app_postgres_executor.GetQueryExecutor(ctx, r.pool)

	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

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
				jsonb_agg(
					DISTINCT jsonb_build_object(
						'id', a.id,
						'name', a.name,
						'description', a.description,
						'avatar_url', a.avatar_url
					)
				) FILTER (WHERE a.id IS NOT NULL),
				'[]'::jsonb
			),
			COALESCE(
				jsonb_agg(
					DISTINCT g.name
				) FILTER (WHERE g.id IS NOT NULL),
				'[]'::jsonb
			)
		FROM track_artists AS selected_artist
		JOIN tracks AS t
			ON t.id = selected_artist.track_id
		JOIN track_statuses AS ts
			ON ts.id = t.status_id
		LEFT JOIN track_artists AS ta
			ON ta.track_id = t.id
		LEFT JOIN artists AS a
			ON a.id = ta.artist_id
		LEFT JOIN track_genres AS tg
			ON tg.track_id = t.id
		LEFT JOIN genres AS g
			ON g.id = tg.genre_id
		WHERE selected_artist.artist_id = $1
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
			ts.code
		ORDER BY
			t.release_date DESC,
			t.created_at DESC
	`

	rows, err := executor.Query(ctx, query, artistId)
	if err != nil {
		return nil, fmt.Errorf("get tracks by artist id: %w", err)
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
			return nil, fmt.Errorf("scan track by artist id: %w", err)
		}

		var genres []string

		if err := json.Unmarshal(trackModel.GenresJSON, &genres); err != nil {
			return nil, fmt.Errorf("unmarshal track genres: %w", err)
		}

		tracks = append(tracks, domain.Track{
			Id:              trackModel.Id,
			AlbumId:         trackModel.AlbumId,
			Name:            trackModel.Name,
			Description:     trackModel.Description,
			CoverStorageKey: trackModel.CoverStorageKey,
			DurationSecond:  time.Duration(trackModel.DurationSecond),
			ReleaseDate:     trackModel.ReleaseDate,
			IsExplicit:      trackModel.IsExplicit,
			IsStreamable:    trackModel.IsStreamable,
			IsDownloadable:  trackModel.IsDownloadable,
			AudioStorageKey: trackModel.AudioStorageKey,
			Status:          trackModel.Status,
			Genres:          genres,
		})
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate tracks by artist id: %w", err)
	}

	return tracks, nil
}
