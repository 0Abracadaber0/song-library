package database

import (
	"database/sql"
	"fmt"
	"log/slog"
	"song_library/internal/config"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations(log *slog.Logger, db *sql.DB, cfg *config.Config) error {
	m, err := migrate.New(
		"file://migrations/",
		getConnectionString(cfg),
	)
	if err != nil {
		return fmt.Errorf("failed to initialize migrate: %w", err)
	}

	log.Info("Applying migrations...")
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to apply migrations: %w", err)
	}
	return nil
}
