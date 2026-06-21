package app_oauth

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type OAuthConfig struct {
	GoogleClientID     string `envconfig:"GOOGLE_CLIENT_ID" required:"true"`
	GoogleClientSecret string `envconfig:"GOOGLE_CLIENT_SECRET" required:"true"`
	GoogleRedirectURL  string `envconfig:"GOOGLE_REDIRECT_URL" required:"true"`
	FrontendURL        string `envconfig:"FRONTEND_URL" required:"true"`
}

func NewOAuthConfig() (OAuthConfig, error) {
	var config OAuthConfig

	if err := envconfig.Process("OAUTH", &config); err != nil {
		return OAuthConfig{}, fmt.Errorf("process envconfig: %w", err)
	}

	return config, nil
}

func NewOAuthConfigMust() OAuthConfig {
	config, err := NewOAuthConfig()

	if err != nil {
		err := fmt.Errorf("get OAuth config: %w", err)
		panic(err)
	}

	return config
}
