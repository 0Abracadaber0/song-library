package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"song_library/internal/config"
	model "song_library/internal/models"
	"time"
)

func GetSong(cfg *config.Config, group, songName string) (model.Song, error) {
	url := "http://" + cfg.ExternalHost.Value + ":" + cfg.ExternalPort.Value

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	uri := url + "/info?group=" + group + "&song=" + songName

	resp, err := client.Get(uri)
	if err != nil {
		return model.Song{}, fmt.Errorf("error sending the request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.Song{}, fmt.Errorf("unexpected status: %s", resp.Status)
	}

	var song model.Song
	if err := json.NewDecoder(resp.Body).Decode(&song); err != nil {
		return model.Song{}, fmt.Errorf("error processing the response: %w", err)
	}

	song.Song = songName
	song.Group = group
	return song, nil
}
