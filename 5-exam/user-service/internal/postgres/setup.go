package postgres

import (
	"database/sql"
	"fmt"
)

type Postgres struct {
	DB *sql.DB
}

func ConnectPostgres(driver, postgres_url string) (*Postgres, error) {
	db, err := sql.Open(driver, postgres_url)
	if err != nil {
		return nil, fmt.Errorf("failed to open sql DB: %v", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed sql DB connection: %v", err)
	}
	return &Postgres{DB: db}, err
}
