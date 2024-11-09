package config

import (
	"log"

	"github.com/SergeyBogomolovv/go-jwt-auth/pkg/env"
	"github.com/joho/godotenv"
)

type Config struct {
	PostgresURL string
	Host        string
	Port        uint16
	JwtSecret   []byte
}

func New() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	return &Config{
		PostgresURL: env.GetString("POSTGRES_URL", "postgres://postgres:Bogomolov980@localhost:5432/go-auth?sslmode=disable"),
		Host:        env.GetString("HOST", "localhost"),
		Port:        uint16(env.GetInt("PORT", 8080)),
		JwtSecret:   []byte(env.GetString("JWT_SECRET", "secret")),
	}
}
