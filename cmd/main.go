package main

import (
	"log/slog"
	"os"

	"go-jwt-auth/internal/app"
	"go-jwt-auth/internal/config"
	"go-jwt-auth/internal/db"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	slog.SetDefault(logger)

	cfg, err := config.New()
	if err != nil {
		slog.Error("Error loading config", "error", err)
		panic(err)
	}

	db, err := db.New(cfg.PostgresURL)
	if err != nil {
		slog.Error("Error connecting to database", "error", err)
		panic(err)
	}
	defer db.Close()
	slog.Info("Connected to database")

	app.New(db, cfg).Run()
}
