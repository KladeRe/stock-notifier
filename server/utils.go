package main

import (
    "fmt"
    "os"
    "encoding/json"
    "github.com/joho/godotenv"
    "log"
)

type Quote struct {
    Global_Quote struct {
        Symbol string `json:"01. symbol"`
        Open string `json:"02. open"`
        High string `json:"03. high"`
        Low string `json:"04. low"`
        Price string `json:"05. price"`
        Volume string `json:"06. volume"`
        Ltd string `json:"07. latest trading day"`
        Previous_close string `json:"08. previous close"`
        Change string `json:"09. change"`
        Change_percent string `json:"10. change percent"`
    } `json:"Global Quote"`

}

func getAPIKey() string {
    // Getting api key from .env file
    envError := godotenv.Load()
    if envError != nil {
        fmt.Printf("Couldn't load .env file")
    }
    api_key := os.Getenv("API_KEY")

    return api_key

}

func decodeJSON(responseBody []uint8) Quote {
    // variable to store JSON data in
    var quote Quote

    unmarshalError := json.Unmarshal([]byte(responseBody), &quote)

    if unmarshalError != nil {
		log.Fatal(unmarshalError)
	}
    return quote

}