package models

type Song struct {
	ID          string `json:"id"`
	Song        string `json:"song"`
	Group       string `json:"group"`
	ReleaseDate string `json:"releaseDate"`
	Text        string `json:"text"`
	Patronymic  string `json:"patronymic"`
}

type Verse struct {
	VerseNumber int    `json:"verseNumber"`
	VerseText   string `json:"verseText"`
}
