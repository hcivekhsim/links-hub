// Package pkg
package pkg

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func NewConnDB() (*sql.DB, error) {
	connStr := os.Getenv("URL_DB")

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("open DB failed: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("ping DB failed: %w", err)
	}

	return db, nil
}
