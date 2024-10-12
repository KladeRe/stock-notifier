package main

import (
	"fmt"
	"log"
	"strconv"

	api_handler "github.com/KladeRe/stock-server/external/alphavantage"

	"github.com/KladeRe/stock-server/internal/utils"
)

func getStockData(config_data string, api_key string) {
	// Iterate over stored data and fetch stock data based on symbol

	// Fetch response from API
	response, responseError := api_handler.SymbolSearch(config_data, api_key)
	if responseError != nil {
		log.Println(responseError)
	}

	// Try decoding the response into a data structure
	decoded, decodeError := api_handler.DecodeResponseJSON(response)
	if decodeError != nil {
		log.Println(decodeError)
	}

	// Check whether data structure is empty
	sanitized, sanitizationError := api_handler.CheckDecodedJSON(decoded, config_data)
	if sanitizationError != nil {
		log.Println(sanitizationError)
	}

	fmt.Printf("%+v\n", sanitized.Global_Quote.Price)

	parsedPrice, _ := strconv.ParseFloat(sanitized.Global_Quote.Price, 32)

	endPrice := float32(parsedPrice)

	fmt.Println(endPrice)

	fmt.Printf("Better luck next time!")

}

func getAllStocks(config_data []string) {
	api_key, key_error := utils.GetEnvVariable("API_KEY")
	if key_error != nil {
		log.Fatal(key_error)
	}
	for i := 0; i < len(config_data); i++ {
		getStockData(config_data[i], api_key)
	}
}
