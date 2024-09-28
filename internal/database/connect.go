package database

import (
	"database/sql"
	"fmt"
	"log/slog"
	"song_library/internal/config"
	"time"

	_ "github.com/lib/pq"
)

func ConnectDB(
	log *slog.Logger,
	cfg *config.Config,
) (*sql.DB, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DbUser.Value,
		cfg.DbPass.Value,
		cfg.DbHost.Value,
		cfg.DbPort.Value,
		cfg.DbName.Value,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed connect to database: %w", err)
	}

	for i := 0; i < 5; i++ {
		err := db.Ping()
		if err == nil {
			break
		}
		log.Info("Failed to connect to database. Retrying...")
		time.Sleep(2 * time.Second)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed ping check: %w", err)
	}

	log.Info("Succesfull connect to database")
	return db, nil
}
