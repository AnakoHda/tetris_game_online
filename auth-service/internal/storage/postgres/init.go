package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"os"
)

func InitPostgres() (*sqlx.DB, error) {
	connStr := os.Getenv("DB_URL")
	if connStr == "" {
		return nil, fmt.Errorf("DB_URL is not set")
	}

	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return db, nil
}
