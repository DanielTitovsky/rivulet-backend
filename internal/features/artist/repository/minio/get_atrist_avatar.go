package artist_minio_repository

import (
	"context"
	"fmt"
	"time"
)

func (r *ArtistStorageRepository) GetArtistAvatar(ctx context.Context, link string) (string, error) {
	expir := time.Hour * 24 * 7

	url, err := r.storage.PresignedGetObject(
		ctx,
		link,
		expir,
	)

	if err != nil {
		return "", fmt.Errorf("get artist avarar link: %w", err)
	}

	return url, nil
}
