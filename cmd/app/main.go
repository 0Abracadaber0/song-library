package main

import (
	"log/slog"
	"os"
	"song_library/internal/config"
	"song_library/internal/database"
	"song_library/internal/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg := config.MustLoad()

	log := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	log.Info("app is starting...", "cfg", cfg)

	app := fiber.New()

	db, err := database.ConnectDB(log, cfg)
	if err != nil {
		panic(err)
	}
	log.Info("succesfull connection to the database")

	if err := database.RunMigrations(log, db, cfg); err != nil {
		panic(err)
	}
	log.Info("succesfull migrations")

	router.SetupRoutes(app)

	if err := app.Listen(cfg.AppHost.Value); err != nil {
		panic("failed start of app " + err.Error())
	}
}
