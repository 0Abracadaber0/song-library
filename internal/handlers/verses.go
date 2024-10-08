package handlers

import (
	"log/slog"
	"song_library/internal/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func VersesHandler(ctx *fiber.Ctx, log *slog.Logger) error {
	songID := ctx.Params("id")

	pageStr := ctx.Query("page", "1")
	limitStr := ctx.Query("limit", "5")

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

	verses, err := service.OutputVerses(songID, limit, offset)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
			"error": err.Error(),
		})
	}

	return ctx.JSON(map[string]interface{}{
		"page":   page,
		"limit":  limit,
		"verses": verses,
	})
}
