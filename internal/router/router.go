package router

import (
	"log/slog"
	"song_library/internal/config"
	handler "song_library/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, cfg *config.Config, log *slog.Logger) {
	// Get list of songs
	app.Get("/songs", handler.SongsHandler)

	// Get verses
	app.Get("/songs/:id/verses", func(ctx *fiber.Ctx) error {
		log.Info("GET /songs/:id/verses")
		return handler.VersesHandler(ctx, log)
	})

	// Delete the song
	app.Delete("/songs/:id", func(ctx *fiber.Ctx) error {
		log.Info("DELETE /songs/:id")
		return handler.DeleteSongHandler(ctx, log)
	})

	// Update song
	app.Put("/songs/:id", func(ctx *fiber.Ctx) error {
		log.Info("PUT /songs/:id")
		return handler.UpdateSongHandler(ctx, cfg, log)
	})

	// Add song
	app.Post("/songs", func(ctx *fiber.Ctx) error {
		log.Info("POST /songs")
		return handler.AddSongHandler(ctx, cfg, log)
	})

}
