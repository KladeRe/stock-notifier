package utils

import (
    "fmt"
    "os"
    "github.com/joho/godotenv"

)

func GetAPIKey() string {
    // Getting api key from .env file
    envError := godotenv.Load()
    if envError != nil {
        fmt.Printf("Couldn't load .env file")
    }
    api_key := os.Getenv("API_KEY")

    return api_key

}

