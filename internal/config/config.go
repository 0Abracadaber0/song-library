package config

import (
	"encoding/json"
	"os"
	"reflect"
)

type Config struct {
	AppHost      StringValue       `env:"APP_HOST" env-default:"localhost:8080"`
	DbPort       StringValue       `env:"POSTGRES_PORT" env-default:"5432"`
	DbHost       StringValue       `env:"POSTGRES_HOST" env-default:"postgres"`
	DbName       StringValue       `env:"POSTGRES_DB" env-default:"library"`
	DbUser       StringValue       `env:"POSTGRES_USER" env-default:"user"`
	DbPass       SecretStringValue `env:"POSTGRES_PASSWORD"`
	ExternalHost StringValue       `env:"EXTERNAL_HOST"`
}

type StringValue struct {
	Value string
}

func (s StringValue) String() string {
	return s.Value
}

func (s StringValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

type SecretStringValue struct {
	Value string
}

func (s SecretStringValue) String() string {
	return "[HIDDEN]"
}

func (s SecretStringValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func MustLoad() *Config {
	var cfg Config

	if err := readConfig(&cfg); err != nil {
		panic("Failed to read env: " + err.Error())
	}

	return &cfg
}

func readConfig(cfg *Config) error {
	val := reflect.ValueOf(cfg).Elem()
	typ := reflect.TypeOf(cfg).Elem()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)

		envKey := fieldType.Tag.Get("env")
		defaultValue := fieldType.Tag.Get("env-default")

		envValue, exists := os.LookupEnv(envKey)
		if !exists {
			envValue = defaultValue
		}

		if field.Kind() == reflect.Struct {
			valueField := field.FieldByName("Value")
			if valueField.IsValid() && valueField.CanSet() {
				valueField.SetString(envValue)
			}
		}
	}

	return nil
}
