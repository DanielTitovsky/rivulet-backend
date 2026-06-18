package app_minIo_storage

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Storage interface {
	PutObject(
		ctx context.Context,
		objectName string,
		reader io.Reader,
		objectSize int64,
		contentType string,
	) (minio.UploadInfo, error)

	RemoveObject(
		ctx context.Context,
		objectName string,
	) error

	PresignedGetObject(
		ctx context.Context,
		objectName string,
		expires time.Duration,
	) (string, error)

	StatObject(
		ctx context.Context,
		objectName string,
	) (minio.ObjectInfo, error)
}

type MinioStorage struct {
	client     *minio.Client
	bucketName string
}

func NewMinioStorage(ctx context.Context, config Config) (*MinioStorage, error) {
	client, err := minio.New(config.MinioEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.MinioRootUser, config.MinioRootPassword, ""),
		Secure: config.MinioUseSSL,
	})

	if err != nil {
		return nil, fmt.Errorf("create minio client: %w", err)
	}

	exists, err := client.BucketExists(ctx, config.BucketName)
	if err != nil {
		return nil, fmt.Errorf("check bucket exists: %w", err)
	}

	if !exists {
		if err := client.MakeBucket(ctx, config.BucketName, minio.MakeBucketOptions{}); err != nil {
			return nil, fmt.Errorf("make bucket: %w", err)
		}
	}

	return &MinioStorage{
		client:     client,
		bucketName: config.BucketName,
	}, nil
}

func (s *MinioStorage) PutObject(
	ctx context.Context,
	objectName string,
	reader io.Reader,
	objectSize int64,
	contentType string,
) (minio.UploadInfo, error) {
	info, err := s.client.PutObject(
		ctx,
		s.bucketName,
		objectName,
		reader,
		objectSize,
		minio.PutObjectOptions{
			ContentType: contentType,
		},
	)

	if err != nil {
		return minio.UploadInfo{}, fmt.Errorf("put object: %w", err)
	}

	return info, nil
}

func (s *MinioStorage) RemoveObject(
	ctx context.Context,
	objectName string,
) error {
	err := s.client.RemoveObject(
		ctx,
		s.bucketName,
		objectName,
		minio.RemoveObjectOptions{},
	)

	if err != nil {
		return fmt.Errorf("remove object: %w", err)
	}

	return nil
}

func (s *MinioStorage) PresignedGetObject(
	ctx context.Context,
	objectName string,
	expires time.Duration,
) (string, error) {
	url, err := s.client.PresignedGetObject(
		ctx,
		s.bucketName,
		objectName,
		expires,
		nil,
	)

	if err != nil {
		return "", fmt.Errorf("presigned get object: %w", err)
	}

	return url.String(), nil
}

func (s *MinioStorage) StatObject(
	ctx context.Context,
	objectName string,
) (minio.ObjectInfo, error) {
	info, err := s.client.StatObject(
		ctx,
		s.bucketName,
		objectName,
		minio.StatObjectOptions{},
	)

	if err != nil {
		return minio.ObjectInfo{}, fmt.Errorf("stat object: %w", err)
	}

	return info, nil
}
