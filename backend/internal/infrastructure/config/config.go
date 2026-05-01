package config

import (
	"fmt"
	"os"
)

type Config struct {
	ServerAddress    string
	FrontendOrigin   string
	JWTSecret        string
	DemoClientID     string
	DemoClientSecret string
}

func Load() (Config, error) {
	config := Config{
		ServerAddress:    envOrDefault("SERVER_ADDRESS", ":18080"),
		FrontendOrigin:   envOrDefault("FRONTEND_ORIGIN", "http://localhost:15173"),
		JWTSecret:        os.Getenv("JWT_SECRET"),
		DemoClientID:     envOrDefault("DEMO_CLIENT_ID", "sezzle-demo-client"),
		DemoClientSecret: os.Getenv("DEMO_CLIENT_SECRET"),
	}

	if config.JWTSecret == "" {
		return Config{}, fmt.Errorf("JWT_SECRET is required")
	}

	if config.DemoClientSecret == "" {
		return Config{}, fmt.Errorf("DEMO_CLIENT_SECRET is required")
	}

	return config, nil
}

func envOrDefault(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return fallback
}
