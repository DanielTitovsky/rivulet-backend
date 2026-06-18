package app_minIo_storage

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port              string `envconfig:"PORT" required:"true"`
	MinioEndpoint     string `envconfig:"ENDPOINT" default:"5432"`
	BucketName        string `envconfig:"BUCKET_NAME" required:"true"`
	MinioRootUser     string `envconfig:"USER" required:"true"`
	MinioRootPassword string `envconfig:"PASSWORD" required:"true"`
	MinioUseSSL       bool   `envconfig:"SSL_MODE" required:"true"`
}

func NewConfig() (*Config, error) {
	var config Config

	if err := envconfig.Process("MINIO", &config); err != nil {
		return &Config{}, fmt.Errorf("Process config: %w", err)
	}

	return &config, nil
}

func NewConfigMust() *Config {
	var config Config

	if err := envconfig.Process("MINIO", &config); err != nil {
		err = fmt.Errorf("Get Minio storage config: %w", err)
		panic(err)
	}

	return &config
}
