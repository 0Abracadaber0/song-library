package main

import (
	"log/slog"
	"os"
	"song_library/internal/config"
	"song_library/internal/database"
	"song_library/internal/router"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg := config.MustLoad()

	log := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	log.Info("app is starting...", "cfg", cfg)

	app := fiber.New(fiber.Config{
		ReadTimeout:  20 * time.Second,
		WriteTimeout: 20 * time.Second,
	})

	if err := database.ConnectDB(log, cfg); err != nil {
		panic(err)
	}
	log.Info("succesfull connection to the database")

	if err := database.RunMigrations(log, cfg); err != nil {
		panic(err)
	}
	log.Info("succesfull migrations")

	router.SetupRoutes(app, cfg, log)

	log.Info("Starting Fiber on", "host", cfg.AppHost.Value)

	if err := app.Listen(cfg.AppHost.Value + ":" + cfg.AppPort.Value); err != nil {
		panic(err)
	}
}
