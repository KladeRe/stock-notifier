package utils

import (
	"os"

	"github.com/joho/godotenv"
)

func GetEnvVariable(name string) (string, error) {
	// Getting api key from .env file
	envError := godotenv.Load()
	if envError != nil {
		return "", envError
	}
	api_key := os.Getenv(name)

	return api_key, nil
}
