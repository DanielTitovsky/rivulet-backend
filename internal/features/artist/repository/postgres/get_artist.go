package artist_postgres_repository

import (
	"context"
	"fmt"
	"time"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
	app_postgres_executor "github.com/DanielTitovsky/rivulet-backend.git/internal/app/repository/postgres"
	"github.com/google/uuid"
)

func (r *ArtistRepository) GetArtist(ctx context.Context, artistId uuid.UUID) (domain.Artist, error) {
	var artistModel ArtistModel

	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	executor := app_postgres_executor.GetQueryExecutor(ctx, r.pool)

	query := `
	SELECT 
	a.id ,
	a."name" ,
	a.description,
	a.avatar_url
	FROM artists a
	WHERE a.id = $1
	`

	err := executor.QueryRow(
		ctx,
		query,
		artistId,
	).Scan(
		&artistModel.Id,
		&artistModel.Name,
		&artistModel.Description,
		&artistModel.AvatarStorageKey,
	)

	if err != nil {
		return domain.Artist{}, fmt.Errorf("Scan artist: %w", err)
	}

	return domain.Artist{
		Id:          artistModel.Id,
		Name:        artistModel.Name,
		Description: artistModel.Description,
		AvatarUrl:   artistModel.AvatarStorageKey,
	}, nil
}
