package helper

import (
	"log"

	"github.com/joho/godotenv"
)

func EnviromentHelper() {

	env := "DEV"
	// env := "PROD"

	if env == "DEV" {
		godotenv.Load(".env.dev")
	} else if env == "PROD" {
		godotenv.Load(".env")
	} else {
		log.Fatalf("Error loading .env file")
	}

}
