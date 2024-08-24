package main

import (
    "net/http"
    "fmt"
    "io"
    "log"
    "github.com/KladeRe/stock-server/config"
)


func main() {
    result := symbolSearch("IBM")
    fmt.Printf("%+v\n", result)

    fileData, err := config.ReadFile("./config.json")
    if (err != nil) {
        log.Fatal("Error while opening file")
        return
    }
    fmt.Printf(string(fileData))
    result2, err2 := config.DecodeJSONConfig(fileData)
    if (err2 != nil) {
        log.Fatal("Error while converting to built in data structure")
    }

    fmt.Printf(result2[0].Notification)
    return 
}


func symbolSearch(keyword string) Quote {

    api_key := getAPIKey()
    
    url := "https://www.alphavantage.co/query?function=GLOBAL_QUOTE&symbol=" + keyword + "&apikey=" +  api_key
    resp, requestErr := http.Get(url)
    if requestErr != nil {
        log.Fatal(requestErr)
    }

    defer resp.Body.Close()
    body, readErr := io.ReadAll(resp.Body)
    if (readErr != nil) {
        log.Fatal(readErr)
    }

    return decodeJSON(body)

}






