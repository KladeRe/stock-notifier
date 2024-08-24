package config

import(
    "encoding/json"
    "io/ioutil"
    "log"

)

type StockConfig struct {
    Symbol string `json:"symbol"`
    Value float32 `json:"value"`
    Buy bool `json:"buy"`
    Notification string `json:"notification"`
}



func DecodeJSONConfig(content []uint8) []StockConfig {
    var config []StockConfig

    unmarshalError := json.Unmarshal([]byte(content), &config)

    if (unmarshalError != nil) {
        log.Fatal(unmarshalError)
    }

    return config
}

func ReadFile(path string) ([]uint8, error) {
    content, err := ioutil.ReadFile(path)
    if err != nil {
        return []uint8{}, err
    }

    return content, nil
}