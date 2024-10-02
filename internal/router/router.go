package router

import (
	"log/slog"
	"song_library/internal/config"
	handler "song_library/internal/handlers"

	_ "song_library/docs" // Swagger docs

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

// @Summary Get all songs with pagination and filtering
// @Description Retrieves a paginated list of songs from the library with optional filtering by song, group, release date, and patronymic
// @Tags songs
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Page limit" default(10)
// @Param song query string false "Filter by song title"
// @Param group query string false "Filter by group name"
// @Param releaseDate query string false "Filter by release date (format: DD.MM.YYYY)"
// @Param patronymic query string false "Filter by patronymic"
// @Success 200 {array} models.Song "List of songs"
// @Failure 400 {object} map[string]interface{} "Invalid query parameters"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /songs [get]
func GetSongsHandler(ctx *fiber.Ctx) error {
	log := ctx.Locals("log").(*slog.Logger)
	log.Info("GET /songs/")
	return handler.SongsHandler(ctx)
}

// @Summary Get song verses
// @Description Retrieves verses of a specific song by ID with pagination
// @Tags songs
// @Produce json
// @Param id path int true "Song ID"
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Page limit" default(10)
// @Success 200 {array} string
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /songs/{id}/verses [get]
func GetSongVersesHandler(ctx *fiber.Ctx) error {
	log := ctx.Locals("log").(*slog.Logger)
	log.Info("GET /songs/:id/verses")
	return handler.VersesHandler(ctx, log)
}

// @Summary Delete song
// @Description Deletes a song by ID
// @Tags songs
// @Produce json
// @Param id path int true "Song ID"
// @Success 200 {string} string "The song has been deleted"
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /songs/{id} [delete]
func DeleteSongHandler(ctx *fiber.Ctx) error {
	log := ctx.Locals("log").(*slog.Logger)
	log.Info("DELETE /songs/:id")
	return handler.DeleteSongHandler(ctx, log)
}

// @Summary Update song
// @Description Updates song details by ID
// @Tags songs
// @Produce json
// @Param id path int true "Song ID"
// @Param song body models.Song true "Song data"
// @Success 200 {string} string "The song has been updated"
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /songs/{id} [put]
func UpdateSongHandler(ctx *fiber.Ctx) error {
	cfg := ctx.Locals("cfg").(*config.Config)
	log := ctx.Locals("log").(*slog.Logger)
	log.Info("PUT /songs/:id")
	return handler.UpdateSongHandler(ctx, cfg, log)
}

// @Summary Add new song
// @Description Adds a new song to the library
// @Tags songs
// @Produce json
// @Param song body models.SongRequest true "Song data"
// @Success 201 {object} models.Song
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /songs [post]
func AddSongHandler(ctx *fiber.Ctx) error {
	cfg := ctx.Locals("cfg").(*config.Config)
	log := ctx.Locals("log").(*slog.Logger)
	log.Info("POST /songs")
	return handler.AddSongHandler(ctx, cfg, log)
}

// SetupRoutes initializes the API routes and Swagger documentation.
func SetupRoutes(app *fiber.App, cfg *config.Config, log *slog.Logger) {
	// Swagger UI route
	app.Get("/swagger/*", swagger.HandlerDefault)

	// Pass logger and config through context
	app.Use(func(ctx *fiber.Ctx) error {
		ctx.Locals("log", log)
		ctx.Locals("cfg", cfg)
		return ctx.Next()
	})

	// Swagger UI маршрут
	app.Get("/swagger/*", swagger.HandlerDefault)
	// Register routes
	app.Get("/songs", GetSongsHandler)
	app.Get("/songs/:id/verses", GetSongVersesHandler)
	app.Delete("/songs/:id", DeleteSongHandler)
	app.Put("/songs/:id", UpdateSongHandler)
	app.Post("/songs", AddSongHandler)
}
