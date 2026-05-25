package app_http_server

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type HttpServerConfig struct {
	Addr         string        `envconfig:"ADDR" required:"true"`
	ShutdownTime time.Duration `envconfig:"SHUTDOWN_TIMEOUT" required:"true"`
}

func NewConfig() (HttpServerConfig, error) {
	var config HttpServerConfig

	if err := envconfig.Process("HTTP_API", &config); err != nil {
		return HttpServerConfig{}, fmt.Errorf("process envconfig: %w", err)
	}

	return config, nil
}

func NewConfigMust() HttpServerConfig {
	config, err := NewConfig()

	if err != nil {
		err := fmt.Errorf("get HttpServer config: %w", err)
		panic(err)
	}

	return config
}
