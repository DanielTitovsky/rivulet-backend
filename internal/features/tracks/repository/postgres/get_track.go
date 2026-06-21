package tracks_postgres_repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
	"github.com/google/uuid"
)

func (r *TrackRepository) GetTrack(ctx context.Context, trackId uuid.UUID) (domain.Track, error) {
	var trackModel TrackModel
	ctx, cancel := context.WithTimeout(ctx, r.pool.GetTimeout())

	defer cancel()

	query := `
WITH selected_track AS (
    SELECT *
    FROM tracks
    WHERE id = $1
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

    COALESCE(genre_data.genres, '[]'::jsonb) AS genres,

    COALESCE(artist_data.artists, '[]'::jsonb) AS artists

FROM selected_track t
LEFT JOIN track_statuses ts ON ts.id = t.status_id

LEFT JOIN LATERAL (
    SELECT jsonb_agg(DISTINCT g.name) AS genres
    FROM track_genres tg
    JOIN genres g ON g.id = tg.genre_id
    WHERE tg.track_id = t.id
) genre_data ON true

LEFT JOIN LATERAL (
    SELECT jsonb_agg(
        DISTINCT jsonb_build_object(
            'id', a.id,
            'name', a.name,
            'description', a.description,
            'avatar_url', a.avatar_url
        )
    ) AS artists
    FROM track_artists ta
    JOIN artists a ON a.id = ta.artist_id
    WHERE ta.track_id = t.id
) artist_data ON true;
`

	err := r.pool.QueryRow(
		ctx,
		query,
		trackId,
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
		&trackModel.GenresJSON,
		&trackModel.ArtistsJSON,
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
		CoverStorageKey: trackModel.CoverStorageKey,
		DurationSecond:  time.Duration(trackModel.DurationSecond) * time.Second,
		ReleaseDate:     trackModel.ReleaseDate,
		IsExplicit:      trackModel.IsExplicit,
		IsStreamable:    trackModel.IsStreamable,
		IsDownloadable:  trackModel.IsDownloadable,
		Genres:          genres,
		Status:          trackModel.Status,
		AudioStorageKey: trackModel.AudioStorageKey,
	}, nil
}
