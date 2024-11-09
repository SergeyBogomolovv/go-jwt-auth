package main

import (
	"log"
	"strconv"

	"go-jwt-auth/internal/app"
	"go-jwt-auth/internal/config"
	"go-jwt-auth/internal/db"
)

func main() {
	cfg := config.New()

	db, err := db.New(cfg.PostgresURL)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()

	app.New(db).Run(cfg.Host + ":" + strconv.Itoa(int(cfg.Port)))
}
