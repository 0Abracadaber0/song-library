package handlers

import (
	"log/slog"
	"song_library/internal/config"
	"song_library/internal/database"
	"song_library/internal/service"

	"github.com/gofiber/fiber/v2"
)

func SongsHandler(ctx *fiber.Ctx) error {
	return ctx.SendString("hi")
}

func LyricsHandler(ctx *fiber.Ctx) error {
	return nil
}

func DeleteSongHandler(ctx *fiber.Ctx, cfg *config.Config, log *slog.Logger) error {
	songID := ctx.Params("id")

	result, err := database.DB.Exec("DELETE FROM songs WHERE id = $1", songID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if rowsAffected == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Song not found",
		})
	}

	log.Info("the song has been deleted")
	return ctx.SendString("The song has been deleted")
}

func UpdateSongHandler(ctx *fiber.Ctx) error {
	return nil
}

type SongRequest struct {
	Group    string `json:"group"`
	SongName string `json:"song"`
}

func AddSongHandler(ctx *fiber.Ctx, cfg *config.Config, log *slog.Logger) error {
	var request SongRequest

	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request format",
		})
	}

	group := request.Group
	songName := request.SongName

	song, err := service.GetSong(cfg, group, songName)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	_, err = database.DB.Exec("INSERT INTO songs"+
		"(song, \"group\", release_date, text, patronymic) "+
		"VALUES ($1, $2, $3, $4, $5)",
		song.Song,
		song.Group,
		song.ReleaseDate,
		song.Text,
		song.Patronymic,
	)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err,
		})
	}

	log.Info("the song has been added")
	return ctx.SendString("The song has been added")
}
