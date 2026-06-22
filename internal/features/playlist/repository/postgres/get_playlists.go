package playlist_postgres_repository

import (
	"context"
	"fmt"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
	app_postgres_executor "github.com/DanielTitovsky/rivulet-backend.git/internal/app/repository/postgres"
)

func (r *PlaylistRepository) GetPlaylists(ctx context.Context) ([]domain.Playlsit, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.GetTimeout())
	defer cancel()

	executor := app_postgres_executor.GetQueryExecutor(ctx, r.pool)

	query := `
SELECT
	p.id,
	p.user_id,
	p.title,
	COALESCE(p.description, ''),
	COALESCE(p.cover_url, ''),
	COALESCE(p.visibility, 'true') = 'true' AS visibility
FROM playlists p
ORDER BY p.created_at DESC;
`

	rows, err := executor.Query(ctx, query)

	if err != nil {
		return nil, fmt.Errorf("Select playlists: %w", err)
	}

	defer rows.Close()

	playlists := make([]domain.Playlsit, 0)

	for rows.Next() {
		var playlistModel PlaylistModel

		err = rows.Scan(
			&playlistModel.Id,
			&playlistModel.UserId,
			&playlistModel.Name,
			&playlistModel.Description,
			&playlistModel.CoverUrl,
			&playlistModel.Visibility,
		)

		if err != nil {
			return nil, fmt.Errorf("Scan playlist: %w", err)
		}

		playlists = append(playlists, domain.Playlsit{
			Id:          playlistModel.Id,
			UserId:      playlistModel.UserId,
			Name:        playlistModel.Name,
			Description: playlistModel.Description,
			CoverUrl:    playlistModel.CoverUrl,
			Visibility:  playlistModel.Visibility,
		})
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("Iterate playlists: %w", err)
	}

	return playlists, nil
}
