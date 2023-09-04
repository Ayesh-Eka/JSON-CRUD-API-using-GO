package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadCnvVariables() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}