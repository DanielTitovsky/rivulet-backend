package playlis_minio_repository

import (
	"context"
	"fmt"
	"strings"
	"time"
)

func (r *PlaylistStorageRepository) GetPlaylistCoverLink(
	ctx context.Context,
	link string,
) (string, error) {
	link = strings.TrimLeft(link, "/")

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
