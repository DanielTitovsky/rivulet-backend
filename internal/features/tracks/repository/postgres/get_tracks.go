package tracks_postgres_repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
	app_postgres_executor "github.com/DanielTitovsky/rivulet-backend.git/internal/app/repository/postgres"
)

func (r *TrackRepository) GetTracks(ctx context.Context, trackFilter domain.TrackFilters) ([]domain.Track, error) {
	trackModels := make([]TrackModel, 0, trackFilter.Limit)
	executor := app_postgres_executor.GetQueryExecutor(ctx, r.pool)

	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)

	defer cancel()

	query := `
WITH random_tracks AS (
    SELECT id 
    FROM tracks t
    WHERE t.status_id = 1
    ORDER BY RANDOM() 
    LIMIT $1
    OFFSET $2
)
SELECT 
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
    ts.name AS status_name,
    COALESCE(
        json_agg(
            json_build_object(
                'id', a.id,
                'name', a.name,
                'description', a.description,
                'avatar_url', a.avatar_url
            )
        ) FILTER (WHERE a.id IS NOT NULL),
        '[]'::json
    ) AS artists,
    COALESCE(
        json_agg(DISTINCT g.name) FILTER (WHERE g.name IS NOT NULL),
        '[]'::json
    ) AS genres
FROM random_tracks rt
JOIN tracks t ON t.id = rt.id
LEFT JOIN track_statuses ts ON ts.id = t.status_id
LEFT JOIN track_genres tg ON tg.track_id = t.id
LEFT JOIN genres g ON g.id = tg.genre_id
LEFT JOIN track_artists ta ON ta.track_id = t.id
LEFT JOIN artists a ON a.id = ta.artist_id
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
    ts.name;
`
	rows, err := executor.Query(
		ctx,
		query,
		trackFilter.Limit,
		trackFilter.Offset,
	)

	if err != nil {
		return []domain.Track{}, fmt.Errorf("Select tracks: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		var trackModel TrackModel

		if err := rows.Scan(
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
		); err != nil {
			return []domain.Track{}, fmt.Errorf("Scan track: %w", err)
		}

		trackModels = append(trackModels, trackModel)
	}

	if err := rows.Err(); err != nil {
		return []domain.Track{}, fmt.Errorf("iterate tracks rows: %w", err)
	}

	trackDomains, err := r.trackModelsToDomain(trackModels)

	if err != nil {
		return []domain.Track{}, fmt.Errorf("Cast trackModels to domain: %w", err)
	}

	return trackDomains, nil
}

func (r *TrackRepository) trackModelsToDomain(models []TrackModel) ([]domain.Track, error) {
	trackDomains := make([]domain.Track, 0, len(models))

	for _, trackModel := range models {
		var artistsModels []trackArtistModel

		if err := json.Unmarshal(trackModel.ArtistsJSON, &artistsModels); err != nil {
			return []domain.Track{}, fmt.Errorf("unmarshal artists: %w", err)
		}

		artists := make([]domain.Artist, 0, len(artistsModels))

		for _, artist := range artistsModels {
			artists = append(artists, domain.Artist{
				Id:          artist.Id,
				Name:        artist.Name,
				Description: artist.Description,
				AvatarUrl:   artist.AvatarStorageKey,
			})
		}

		var genres []string

		if err := json.Unmarshal(trackModel.GenresJSON, &genres); err != nil {
			return []domain.Track{}, fmt.Errorf("unmarshal genres: %w", err)
		}

		trackDomains = append(trackDomains, domain.Track{
			Id:              trackModel.Id,
			Name:            trackModel.Name,
			Description:     trackModel.Description,
			AlbumId:         trackModel.AlbumId,
			Artists:         artists,
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
	return trackDomains, nil
}
