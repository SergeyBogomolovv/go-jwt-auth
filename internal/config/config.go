package config

import (
	"fmt"

	"go-jwt-auth/pkg/env"

	"github.com/joho/godotenv"
)

type Config struct {
	PostgresURL string
	Host        string
	Port        uint16
	JwtSecret   []byte
}

func New() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}

	return &Config{
		PostgresURL: env.GetString("POSTGRES_URL", "postgres://postgres:Bogomolov980@localhost:5432/go-auth?sslmode=disable"),
		Host:        env.GetString("HOST", "localhost"),
		Port:        uint16(env.GetInt("PORT", 8080)),
		JwtSecret:   []byte(env.GetString("JWT_SECRET", "secret")),
	}, nil
}
