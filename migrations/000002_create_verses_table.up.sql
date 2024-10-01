CREATE TABLE IF NOT EXISTS verses (
    id SERIAL PRIMARY KEY,
    song_id INT NOT NULL,
    verse_number INT NOT NULL,
    verse_text TEXT NOT NULL,
    FOREIGN KEY (song_id) REFERENCES songs (id) ON DELETE CASCADE
);