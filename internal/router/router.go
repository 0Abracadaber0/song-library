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

	// Get lyrics
	app.Get("/songs/:id/lyrics", handler.LyricsHandler)

	// Delete the song
	app.Delete("/songs/:id", handler.DeleteSongHandler)

	// Update song
	app.Put("/songs/:id", handler.UpdateSongHandler)

	// Add song
	app.Post("/songs", func(ctx *fiber.Ctx) error {
		log.Info("POST /songs")
		return handler.AddSongHandler(ctx, cfg)
	})

}
