package service

import (
	"fmt"
	"song_library/internal/database"
	model "song_library/internal/models"
	"strings"
)

func SplitTextIntoVerses(text string) []model.Verse {
	versesText := strings.Split(text, "\n\n")

	var verses []model.Verse
	for i, verseText := range versesText {
		verse := model.Verse{
			VerseNumber: i + 1,
			VerseText:   verseText,
		}
		verses = append(verses, verse)
	}
	return verses
}

func SaveVerses(songID int, verses []model.Verse) error {
	for _, verse := range verses {
		_, err := database.DB.Exec(
			`INSERT INTO verses (song_id, verse_number, verse_text) VALUES ($1, $2, $3)`,
			songID,
			verse.VerseNumber,
			verse.VerseText,
		)
		if err != nil {
			return fmt.Errorf("error saving verse: %w", err)
		}
	}
	return nil
}

func OutputVerses(songID string, limit, offset int) ([]string, error) {
	rows, err := database.DB.Query(
		`SELECT verse_text FROM verses WHERE song_id = $1 ORDER BY id LIMIT $2 OFFSET $3`,
		songID, limit, offset,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve verses: %w", err)
	}
	defer rows.Close()

	var verses []string
	for rows.Next() {
		var verse string
		if err := rows.Scan(&verse); err != nil {
			return nil, fmt.Errorf("failed to scan verse: %w", err)
		}
		verses = append(verses, verse)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during rows iteration: %w", err)
	}

	return verses, nil
}

func VerseToText(songID string) (string, error) {
	rows, err := database.DB.Query(
		`SELECT verse_text FROM verses WHERE song_id = $1`,
		songID,
	)
	if err != nil {
		return "", fmt.Errorf("failed to retrieve verses: %w", err)
	}
	defer rows.Close()

	var verses string
	for rows.Next() {
		var verse string
		if err := rows.Scan(&verse); err != nil {
			return "", fmt.Errorf("failed to scan verse: %w", err)
		}
		if verses != "" {
			verses += "\n\n"
		}
		verses += verse
	}

	if err := rows.Err(); err != nil {
		return "", fmt.Errorf("error during rows iteration: %w", err)
	}

	return verses, nil
}
