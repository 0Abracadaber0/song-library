package main

import (
	"log/slog"
	"os"
	"song_library/internal/config"
)

func main() {
	cfg := config.MustLoad()

	log := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	log.Info("App has been started", "cfg", cfg)
}
