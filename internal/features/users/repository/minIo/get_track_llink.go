package users_minio_repository

import (
	"context"
	"fmt"
	"time"
)

func (r *UserStorageRepository) GetTrackAudioLink(
	ctx context.Context,
	link string,
) (string, error) {
	url, err := r.storage.PresignedGetObject(
		ctx,
		link,
		15*time.Minute,
	)

	if err != nil {
		return "", fmt.Errorf("get track cover link: %w", err)
	}

	return url, nil
}
