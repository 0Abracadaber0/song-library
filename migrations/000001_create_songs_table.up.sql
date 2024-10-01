CREATE TABLE IF NOT EXISTS songs (
    id SERIAL PRIMARY KEY,
    song VARCHAR(255) NOT NULL,
    "group" VARCHAR(255) NOT NULL,
    release_date VARCHAR(20) NOT NULL,
    patronymic VARCHAR(255)
);
