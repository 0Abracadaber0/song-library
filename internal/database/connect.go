package database

import (
	"database/sql"
	"fmt"
	"log/slog"
	"song_library/internal/config"

	_ "github.com/lib/pq"
)

func ConnectDB(
	log *slog.Logger,
	cfg *config.Config,
) (*sql.DB, error) {
	connStr := getConnectionString(cfg)
	log.Info("Connecting with connection string:", connStr, connStr)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed connect to database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed ping check: %w", err)
	}

	log.Info("Succesfull connect to database")
	return db, nil
}

func getConnectionString(cfg *config.Config) string {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DbUser.Value,
		cfg.DbPass.Value,
		cfg.DbHost.Value,
		cfg.DbPort.Value,
		cfg.DbName.Value,
	)

	return connStr
}
