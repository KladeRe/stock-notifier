package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/KladeRe/stock-server/api_handler"
	"github.com/KladeRe/stock-server/config"
	"github.com/KladeRe/stock-server/internal/utils"
)

func getFileData(location string) []config.StockConfig {

	// Get the raw data from the config file
	rawFileData, fileReadErr := config.ReadFile(location)
	if fileReadErr != nil {
		log.Fatal(fileReadErr)
		return []config.StockConfig{}
	}

	// Convert the data in the config to native data structure
	config_data, conversionErr := config.DecodeJSONConfig(rawFileData)
	if conversionErr != nil {
		log.Fatal(conversionErr)
		return []config.StockConfig{}
	}

	return config_data

}

func getStockData(config_data config.StockConfig, api_key string) {
	// Iterate over stored data and fetch stock data based on symbol

	// Fetch response from API
	response, responseError := api_handler.SymbolSearch(config_data.Symbol, api_key)
	if responseError != nil {
		log.Println(responseError)
	}

	// Try decoding the response into a data structure
	decoded, decodeError := api_handler.DecodeResponseJSON(response)
	if decodeError != nil {
		log.Println(decodeError)
	}

	// Check whether data structure is empty
	sanitized, sanitizationError := api_handler.CheckDecodedJSON(decoded, config_data.Symbol)
	if sanitizationError != nil {
		log.Println(sanitizationError)
	}

	fmt.Printf("%+v\n", sanitized.Global_Quote.Price)

	parsedPrice, _ := strconv.ParseFloat(sanitized.Global_Quote.Price, 32)

	endPrice := float32(parsedPrice)

	if config_data.Buy && endPrice <= config_data.Value {
		fmt.Printf(`%s has now dropped below the price you are waiting for`, config_data.Symbol)
		return
	}

	if !config_data.Buy && endPrice >= config_data.Value {
		fmt.Printf(`%s has now risen over the price you are waiting for`, config_data.Symbol)
		return
	}

	fmt.Printf("Better luck next time!")

}

func getAllStocks(config_data []config.StockConfig) {
	api_key, key_error := utils.GetAPIKey()
	if key_error != nil {
		log.Fatal(key_error)
	}
	for i := 0; i < len(config_data); i++ {
		getStockData(config_data[i], api_key)
	}
}
