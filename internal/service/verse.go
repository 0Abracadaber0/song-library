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
