package main

import (
	"fmt"
	"song_library/internal/config"
)

func main() {
	cfg := config.MustLoad()

	// TODO: запретить вывод пароля в логи

	fmt.Println(cfg)
}
