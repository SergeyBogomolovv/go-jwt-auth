package main

import (
	"fmt"
	"log"

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

	fmt.Println(cfg)
}
