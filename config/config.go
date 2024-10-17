package config

import (
	"os"
)

type Config struct {
	DbUser string
	DbPass string
	DbHost string
	DbPort string
	DbName string
}

var ENV *Config

func LoadConfig() {

	ENV = &Config{
		DbUser: os.Getenv("DB_USER"),
		DbPass: os.Getenv("DB_PASSWORD"),
		DbHost: os.Getenv("DB_HOST"),
		DbPort: os.Getenv("DB_PORT"),
		DbName: os.Getenv("DB_NAME"),
	}
}
