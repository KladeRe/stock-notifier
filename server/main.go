package main

import (
    "fmt"
    "log"
    "github.com/KladeRe/stock-server/config"
    "github.com/KladeRe/stock-server/api_handler"
    "github.com/KladeRe/stock-server/utils"
)

func main() {
    

    // Get the raw data from the config file
    rawFileData, fileReadErr := config.ReadFile("./config.json")
    if (fileReadErr!= nil) {
        log.Fatal(fileReadErr)
        return
    }

    fmt.Printf(string(rawFileData))

    // Convert the data in the config to native data structure
    data, conversionErr := config.DecodeJSONConfig(rawFileData)
    if (conversionErr != nil) {
        log.Fatal(conversionErr)
        return
    }

    // Read api key

    api_key, keyError := utils.GetAPIKey()

	if (keyError != nil) {
		log.Fatal(keyError)
	}

    // Iterate over stored data and fetch stock data based on symbol
    for i := 0; i < len(data); i++ {

        // Fetch response from API
        response, responseError := api_handler.SymbolSearch(data[i].Symbol, api_key)
        if (responseError != nil) {
            log.Println(responseError)
        }

        // Try decoding the response into a data structure
        decoded, decodeError := api_handler.DecodeResponseJSON(response)
        if (decodeError != nil) {
            log.Println(decodeError)
        }

        // Check whether data structure is empty
        sanitized, sanitizationError := api_handler.CheckDecodedJSON(decoded, data[i].Symbol)
        if (sanitizationError != nil) {
            log.Println(sanitizationError)
        }

        fmt.Printf("%+v\n", sanitized.Global_Quote.Price)
    }
    
    
    return 
}









