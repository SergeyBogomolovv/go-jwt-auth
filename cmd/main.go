package main

import (
	"log"

	"go-jwt-auth/internal/app"
	"go-jwt-auth/internal/config"
	"go-jwt-auth/internal/db"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	db, err := db.New(cfg.PostgresURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	app.New(db, cfg).Run()
}
