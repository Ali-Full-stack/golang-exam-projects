package storage

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	url := os.Getenv("POSTGRES_URL")

	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, fmt.Errorf("couldn't open database: %w", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("failed connection to database: %w", err)
	}

	return db, nil
}
