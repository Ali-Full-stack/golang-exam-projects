package postgres

import (
	"database/sql"
	"fmt"
)

func ConnectPostgres(postgres_url string) (*sql.DB, error) {
	db, err := sql.Open("postgres", postgres_url)
	if err != nil {
		return nil, fmt.Errorf("failed to open to postgres: %v", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping to postgres: %v", err)
	}
	return db, nil
}
