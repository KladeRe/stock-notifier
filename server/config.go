package main

import(
    "encoding/json"
    "io/ioutil"
    "log"
)

type Config struct {
    Notification string `json:"notification"`
    Stocks []string `json:"stocks"`

}

func decodeJSONConfig(content []uint8) Config {
    var config Config

    unmarshalError := json.Unmarshal([]byte(content), &config)

    if (unmarshalError != nil) {
        log.Fatal(unmarshalError)
    }

    return config
}

func readFile(directory string) []uint8 {
    content, err := ioutil.ReadFile(directory)
    if err != nil {
        log.Fatal("Errow when opening file: ", err)
    }

    return content
}