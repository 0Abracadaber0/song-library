package router

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app *fiber.App) {
	// Get list of songs
	app.Get("/songs")

	// Get lyrics
	app.Get("/songs/:id/lyrics")

	// Delete the song
	app.Delete("/songs/:id")

	// Update song
	app.Put("/songs/:id")

	// Delete song
	app.Post("/songs")
}
