package helper

import (
	"github.com/joho/godotenv"
)

func EnviromentHelper() {
	err := godotenv.Load(".env")
	if err != nil {
		return
	}
}
