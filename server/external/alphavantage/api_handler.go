package api_handler

import (
    "encoding/json"
	"net/http"
	"io"
	"errors"
)

type Quote struct {
    Global_Quote Global_Quote `json:"Global Quote"`

}

type Global_Quote struct {
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
} 


func DecodeResponseJSON(responseBody []uint8) (Quote, error) {
    // variable to store JSON data in
    var quote Quote

    unmarshalError := json.Unmarshal([]byte(responseBody), &quote)

    if unmarshalError != nil {
		return Quote{}, unmarshalError
	}
	
    return quote, nil

}

func SymbolSearch(keyword string, api_key string) ([]uint8, error) {

    url := "https://www.alphavantage.co/query?function=GLOBAL_QUOTE&symbol=" + keyword + "&apikey=" +  api_key
    resp, requestErr := http.Get(url)

    if requestErr != nil {
        return []uint8{}, requestErr
    }

    defer resp.Body.Close()

    body, readErr := io.ReadAll(resp.Body)

	return body, readErr
	
	

}

func CheckDecodedJSON(parsed Quote, keyword string) (Quote, error) {

	// For checking whether the body actually includes any info
	if (parsed == Quote{}) {
		symbolError := errors.New("Couldn't find stock info for " + keyword)
		return Quote{}, symbolError
	}

	return parsed, nil

}