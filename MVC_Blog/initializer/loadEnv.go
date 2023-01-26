package initializer

import (
	"log"
	"github.com/joho/godotenv"
)

func LoadEnvVariable() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}