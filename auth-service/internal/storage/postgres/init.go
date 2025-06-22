package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log/slog"
	"os"
)

func InitPostgres() (*sqlx.DB, error) {
	connStr := os.Getenv("DB_URL")
	if connStr == "" {
		slog.Error("InitPostgres: DB_URL environment variable is not set")
		return nil, fmt.Errorf("DB_URL is not set")
	}

	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		slog.Error("InitPostgres: failed to connect to database", "error", err)
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	slog.Info("InitPostgres: database connection successfully")
	return db, nil
}
