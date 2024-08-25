package utils

import (
    "os"
    "github.com/joho/godotenv"

)

func GetAPIKey() (string, error) {
    // Getting api key from .env file
    envError := godotenv.Load()
    if envError != nil {
        return "", envError
    }
    api_key := os.Getenv("API_KEY")

    return api_key, nil

}

