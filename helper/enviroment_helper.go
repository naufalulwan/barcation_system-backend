package helper

import (
	"github.com/joho/godotenv"
)

func EnviromentHelper() {
	err := godotenv.Load(".env.dev")
	if err != nil {
		return
	}
}
