package playlist_postgres_repository

import (
	"context"
	"fmt"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
	app_postgres_executor "github.com/DanielTitovsky/rivulet-backend.git/internal/app/repository/postgres"
	"github.com/google/uuid"
)

func (r *PlaylistRepository) GetPlaylist(ctx context.Context, playlistId uuid.UUID) (domain.Playlsit, error) {
	var playlistModel PlaylistModel

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
WHERE p.id = $1
LIMIT 1;
`

	err := executor.QueryRow(
		ctx,
		query,
		playlistId,
	).Scan(
		&playlistModel.Id,
		&playlistModel.UserId,
		&playlistModel.Name,
		&playlistModel.Description,
		&playlistModel.CoverUrl,
		&playlistModel.Visibility,
	)

	if err != nil {
		return domain.Playlsit{}, fmt.Errorf("Scan playlist: %w", err)
	}

	return domain.Playlsit{
		Id:          playlistModel.Id,
		UserId:      playlistModel.UserId,
		Name:        playlistModel.Name,
		Description: playlistModel.Description,
		CoverUrl:    playlistModel.CoverUrl,
		Visibility:  playlistModel.Visibility,
	}, nil
}
