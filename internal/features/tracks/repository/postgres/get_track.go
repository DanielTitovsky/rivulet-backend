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
    g.name AS genre,
    COALESCE(
        json_agg(
            json_build_object(
                'id', a.id,
                'name', a.name
            )
        ) FILTER (WHERE a.id IS NOT NULL),
        '[]'
    ) AS artists
FROM tracks t
LEFT JOIN track_statuses ts ON ts.id = t.status_id
LEFT JOIN track_genres tg ON tg.track_id = t.id
LEFT JOIN genres g ON g.id = tg.genre_id
LEFT JOIN track_artists ta ON ta.track_id = t.id
LEFT JOIN artists a ON a.id = ta.artist_id
WHERE t.id = $1
GROUP BY 
    t.id,
    ts.name,
    g.name
LIMIT 1;
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
		&trackModel.Genre,
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
			Id:               artist.Id,
			Name:             artist.Name,
			Description:      artist.Description,
			AvatarStorageKey: artist.AvatarStorageKey,
		})
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
		Status:          trackModel.Status,
		AudioStorageKey: trackModel.AudioStorageKey,
	}, nil
}
