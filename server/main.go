package main

import (
    "net/http"
    "fmt"
    "io"
    "log"
)


func main() {
    result := symbolSearch("IBM")
    fmt.Printf("%+v\n", result)

    result2 := decodeJSONConfig(readFile("./config.json"))

    fmt.Printf(result2.Notification)
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






