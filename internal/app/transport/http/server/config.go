package app_http_server

import (
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
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

func GetCORSConfig() cors.Config {
	return cors.Config{
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:3000" ||
				origin == "http://localhost:3001" ||
				origin == "http://127.0.0.1:3000" ||
				origin == "http://127.0.0.1:3001"
		},

		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"PATCH",
			"DELETE",
			"OPTIONS",
		},

		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Accept",
			"Authorization",
			"X-Requested-With",
			"X-Request-Id",
			"Cache-Control",
			"Pragma",
		},

		ExposeHeaders: []string{
			"Content-Length",
			"X-Request-Id",
		},

		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
}
