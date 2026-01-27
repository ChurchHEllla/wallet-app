package config

import (
	"errors"
	"fmt"
	"os"
)

type Config struct {
	PostgresDSN string
}

func Load() (*Config, error) {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
	)
	if dsn == "" {
		return nil, errors.New("connection configuration is empty")
	}
	fmt.Print(dsn)
	return &Config{
		PostgresDSN: dsn,
	}, nil
}
