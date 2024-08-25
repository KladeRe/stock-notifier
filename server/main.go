package main

import (
    "fmt"
    "log"
    "github.com/KladeRe/stock-server/config"
    "github.com/KladeRe/stock-server/api_handler"
)

func main() {
    rawFileData, fileReadErr := config.ReadFile("./config.json")
    if (fileReadErr!= nil) {
        log.Fatal(fileReadErr)
        return
    }

    fmt.Printf(string(rawFileData))

    data, conversionErr := config.DecodeJSONConfig(rawFileData)
    if (conversionErr != nil) {
        log.Fatal(conversionErr)
        return
    }

    fmt.Printf(data[0].Notification)

    for i := 0; i < len(data); i++ {
        result, searchErr := api_handler.SymbolSearch(data[i].Symbol)
        if (searchErr != nil) {
            log.Fatal(searchErr)
        }
        fmt.Printf("%+v\n", result.Global_Quote.Price)
    }
    
    
    return 
}









