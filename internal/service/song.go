package service

import (
	"fmt"
	"song_library/internal/config"
	"song_library/internal/database"
	model "song_library/internal/models"
)

func AddSong(song *model.Song) error {
	var songID int
	err := database.DB.QueryRow(
		`INSERT INTO songs (song, "group", release_date, patronymic) 
        VALUES ($1, $2, $3, $4) RETURNING id`,
		song.Song, song.Group, song.ReleaseDate, song.Patronymic,
	).Scan(&songID)

	if err != nil {
		return fmt.Errorf("failed to save song: %w", err)
	}

	verses := SplitTextIntoVerses(song.Text)
	if err := SaveVerses(songID, verses); err != nil {
		return fmt.Errorf("failed to save verses: %w", err)
	}

	return nil
}

func UpdateSongWithVerses(cfg *config.Config, songID string, request model.Song) error {
	tx, err := database.DB.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	_, err = tx.Exec(
		`UPDATE songs SET "group" = $1, song = $2, release_date = $3, patronymic = $4 
		WHERE id = $5`,
		request.Group,
		request.Song,
		request.ReleaseDate,
		request.Patronymic,
		songID,
	)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to update song: %w", err)
	}

	_, err = tx.Exec(`DELETE FROM verses WHERE song_id = $1`, songID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to delete old verses: %w", err)
	}

	if request.Text != "" {
		verses := SplitTextIntoVerses(request.Text)

		for _, verse := range verses {
			_, err := tx.Exec(
				`INSERT INTO verses (song_id, verse_number, verse_text) VALUES ($1, $2, $3)`,
				songID,
				verse.VerseNumber,
				verse.VerseText,
			)

			if err != nil {
				tx.Rollback()
				return fmt.Errorf("failed to insert verse: %w", err)
			}
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func DeleteSong(songID string) error {
	tx, err := database.DB.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	_, err = tx.Exec("DELETE FROM verses WHERE song_id = $1", songID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to delete verses: %w", err)
	}

	result, err := tx.Exec("DELETE FROM songs WHERE id = $1", songID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to delete song: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to retrieve rows affected: %w", err)
	}

	if rowsAffected == 0 {
		tx.Rollback()
		return fmt.Errorf("song not found")
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func OutputSongs(limit, offset int) ([]model.Song, error) {
	rows, err := database.DB.Query(
		`SELECT id, song, "group", release_date, patronymic FROM songs ORDER BY id LIMIT $1 OFFSET $2`,
		limit, offset,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve songs: %w", err)
	}
	defer rows.Close()

	var songs []model.Song
	for rows.Next() {
		var song model.Song
		if err := rows.Scan(&song.ID, &song.Song, &song.Group, &song.ReleaseDate, &song.Patronymic); err != nil {
			return nil, fmt.Errorf("failed to scan song: %w", err)
		}
		verses, err := VerseToText(song.ID)
		if err != nil {
			return nil, err
		}
		song.Text = verses
		songs = append(songs, song)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during rows iteration: %w", err)
	}

	return songs, nil
}
