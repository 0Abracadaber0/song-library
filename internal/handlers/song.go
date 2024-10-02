package handlers

import (
	"log/slog"
	"song_library/internal/config"
	model "song_library/internal/models"
	"song_library/internal/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func SongsHandler(ctx *fiber.Ctx) error {
	pageStr := ctx.Query("page", "1")
	limitStr := ctx.Query("limit", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		return ctx.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"error": "Invalid page number",
		})
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		return ctx.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"error": "Invalid limit number",
		})
	}

	offset := (page - 1) * limit

	// Получаем параметры фильтрации из запроса
	song := ctx.Query("song")
	group := ctx.Query("group")
	releaseDate := ctx.Query("releaseDate")
	patronymic := ctx.Query("patronymic")

	// Передаем параметры фильтрации в сервис
	songs, err := service.OutputSongs(song, group, releaseDate, patronymic, limit, offset)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
			"error": "Failed to retrieve songs: " + err.Error(),
		})
	}

	return ctx.JSON(map[string]interface{}{
		"page":  page,
		"limit": limit,
		"songs": songs,
	})
}

func DeleteSongHandler(ctx *fiber.Ctx, log *slog.Logger) error {
	songID := ctx.Params("id")

	if err := service.DeleteSong(songID); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
			"error": err.Error(),
		})
	}

	log.Info("The song has been deleted")
	return ctx.SendString("The song has been deleted")
}

func UpdateSongHandler(ctx *fiber.Ctx, cfg *config.Config, log *slog.Logger) error {
	songID := ctx.Params("id")

	var request model.Song
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"error": "Invalid request format",
		})
	}

	if err := service.UpdateSongWithVerses(cfg, songID, request); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
			"error": err.Error(),
		})
	}

	log.Info("The song has been updated")
	return ctx.SendString("The song has been updated")
}

func AddSongHandler(ctx *fiber.Ctx, cfg *config.Config, log *slog.Logger) error {
	var request model.Song

	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"error": "Invalid request format",
		})
	}

	song, err := service.GetSong(cfg, request.Group, request.Song)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"error": err.Error(),
		})
	}

	if err := service.AddSong(&song); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
			"error": err.Error(),
		})
	}

	log.Info("The song has been added")
	return ctx.SendString("The song has been added")
}
