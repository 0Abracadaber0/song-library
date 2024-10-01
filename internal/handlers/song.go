package handlers

import (
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

func DeleteSongHandler(ctx *fiber.Ctx) error {
	return nil
}

func UpdateSongHandler(ctx *fiber.Ctx) error {
	return nil
}

type SongRequest struct {
	Group    string `json:"group"`
	SongName string `json:"song"`
}

func AddSongHandler(ctx *fiber.Ctx, cfg *config.Config) error {
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

	return ctx.SendString("The song has been added")
}
