package tracks_minio_repository

import (
	"context"
	"fmt"
	"time"
)

func (r *TrackStorageRepository) GetTrackLink(
	ctx context.Context,
	link string,
) (string, error) {
	url, err := r.storage.PresignedGetObject(
		ctx,
		link,
		15*time.Minute,
	)

	if err != nil {
		return "", fmt.Errorf("get track link: %w", err)
	}

	return url, nil
}
