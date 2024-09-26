package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	_ "github.com/joho/godotenv/autoload"
)

type ConfigDatabase struct {
	Port string `env:"DB_PORT" env-default:"5432"`
	Host string `env:"DB_HOST" env-default:"localhost"`
	Name string `env:"DB_NAME" env-default:"postgres"`
	User string `env:"DB_USER" env-default:"user"`
	Pass string `env:"DB_PASS"`
}

func MustLoad() *ConfigDatabase {
	var cfg ConfigDatabase

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		panic("Failed to read env " + err.Error())
	}

	return &cfg
}
