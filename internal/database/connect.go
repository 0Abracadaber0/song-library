package database

import (
	"database/sql"
	"fmt"
	"log/slog"
	"song_library/internal/config"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB(log *slog.Logger, cfg *config.Config) error {
	connStr := getConnectionString(cfg)
	log.Info("Connecting with connection string:", connStr, connStr)

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("failed connect to database: %w", err)
	}

	if err := DB.Ping(); err != nil {
		return fmt.Errorf("failed ping check: %w", err)
	}

	log.Info("Succesfull connect to database")
	return nil
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
