package config

import (
	"fmt"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type Config struct {
	DBSource   string `env:"DB_SOURCE,required"`
	ServerPort string `env:"SERVER_PORT,required"`
}

func NewConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	cfg := new(Config)
	err = env.Parse(cfg)
	if err != nil {
		return nil, fmt.Errorf("error parsing environment variables: %w", err)
	}

	return cfg, nil
}
