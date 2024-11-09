package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func New(uri string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", uri)
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error ping database: %v", err)
	}

	return db, nil
}
