package artist_postgres_repository

import (
	"context"
	"fmt"
	"time"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
	app_postgres_executor "github.com/DanielTitovsky/rivulet-backend.git/internal/app/repository/postgres"
	"github.com/google/uuid"
)

func (r *ArtistRepository) GetArtistAlbums(
	ctx context.Context,
	artistId uuid.UUID,
) ([]domain.Album, error) {
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	executor := app_postgres_executor.GetQueryExecutor(ctx, r.pool)

	query := `
		SELECT
			al.id,
			al.title,
			COALESCE(al.description, ''),
			COALESCE(al.cover_url, ''),
			COALESCE(al.release_date, '1970-01-01'::date)
		FROM album_artists aa
		JOIN albums al
			ON al.id = aa.album_id
		WHERE aa.artist_id = $1
		ORDER BY
			al.release_date DESC,
			al.created_at DESC
	`

	rows, err := executor.Query(
		ctx,
		query,
		artistId,
	)

	if err != nil {
		return nil, fmt.Errorf("Select artist albums: %w", err)
	}

	defer rows.Close()

	albums := make([]domain.Album, 0)

	for rows.Next() {
		var albumModel AlbumModel

		err = rows.Scan(
			&albumModel.Id,
			&albumModel.Title,
			&albumModel.Description,
			&albumModel.CoverStorageKey,
			&albumModel.ReleaseDate,
		)

		if err != nil {
			return nil, fmt.Errorf("Scan artist album: %w", err)
		}

		albums = append(albums, domain.Album{
			Id:           albumModel.Id,
			Name:         albumModel.Title,
			Description:  albumModel.Description,
			CoverUrl:     albumModel.CoverStorageKey,
			Release_date: albumModel.ReleaseDate,
		})
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("Iterate artist albums: %w", err)
	}

	return albums, nil
}
