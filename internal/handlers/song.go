package handlers

import (
	"log/slog"
	"song_library/internal/config"
	model "song_library/internal/models"
	"song_library/internal/service"

	"github.com/gofiber/fiber/v2"
)

func SongsHandler(ctx *fiber.Ctx) error {
	return ctx.SendString("hi")
}

func DeleteSongHandler(ctx *fiber.Ctx, log *slog.Logger) error {
	songID := ctx.Params("id")

	if err := service.DeleteSong(songID); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	log.Info("the song has been deleted")
	return ctx.SendString("The song has been deleted")
}

func UpdateSongHandler(ctx *fiber.Ctx, cfg *config.Config, log *slog.Logger) error {
	songID := ctx.Params("id")

	var request model.Song
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request format",
		})
	}

	if err := service.UpdateSongWithVerses(cfg, songID, request); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	log.Info("The song has been updated")
	return ctx.SendString("The song has been updated")
}

func AddSongHandler(ctx *fiber.Ctx, cfg *config.Config, log *slog.Logger) error {
	var request model.Song

	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request format",
		})
	}

	song, err := service.GetSong(cfg, request.Group, request.Song)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := service.AddSong(&song); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	log.Info("the song has been added")
	return ctx.SendString("The song has been added")
}
