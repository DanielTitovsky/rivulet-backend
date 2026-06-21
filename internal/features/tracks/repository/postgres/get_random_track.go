package tracks_postgres_repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
)

func (r *TrackRepository) GetRandomTrack(ctx context.Context) (domain.Track, error) {
	var trackModel TrackModel
	ctx, cancel := context.WithTimeout(ctx, r.pool.GetTimeout())

	defer cancel()

	query := `
WITH random_track AS (
    SELECT id
    FROM tracks
    ORDER BY RANDOM()
    LIMIT 1
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
        jsonb_agg(
            DISTINCT jsonb_build_object(
                'id', a.id,
                'name', a.name,
    			'description', a.description,
    			'avatar_url', a.avatar_url
            )
        ) FILTER (WHERE a.id IS NOT NULL),
        '[]'::jsonb
    ) AS artists,
    COALESCE(
        jsonb_agg(DISTINCT g.name) FILTER (WHERE g.name IS NOT NULL),
        '[]'::jsonb
    ) AS genres
FROM random_track rt
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

	err := r.pool.QueryRow(
		ctx,
		query,
	).Scan(
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
		return domain.Track{}, fmt.Errorf("Scan track: %w", err)
	}

	var artistsModels []trackArtistModel

	if err := json.Unmarshal(trackModel.ArtistsJSON, &artistsModels); err != nil {
		return domain.Track{}, fmt.Errorf("unmarshal artists: %w", err)
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
		return domain.Track{}, fmt.Errorf("unmarshal genres: %w", err)
	}

	return domain.Track{
		Id:              trackModel.Id,
		Name:            trackModel.Name,
		Description:     trackModel.Description,
		AlbumId:         trackModel.AlbumId,
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
	}, nil
}
