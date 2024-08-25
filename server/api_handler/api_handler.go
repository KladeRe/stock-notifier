package api_handler

import (
    "encoding/json"
	"github.com/KladeRe/stock-server/utils"
	"net/http"
	"io"
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


func DecodeResponseJSON(responseBody []uint8) (Quote, error) {
    // variable to store JSON data in
    var quote Quote

    unmarshalError := json.Unmarshal([]byte(responseBody), &quote)

    if unmarshalError != nil {
		return Quote{}, unmarshalError
	}
    return quote, nil

}

func SymbolSearch(keyword string) (Quote, error) {

    api_key := utils.GetAPIKey()
    
    url := "https://www.alphavantage.co/query?function=GLOBAL_QUOTE&symbol=" + keyword + "&apikey=" +  api_key
    resp, requestErr := http.Get(url)
    if requestErr != nil {
        return Quote{}, requestErr
    }

    defer resp.Body.Close()
    body, readErr := io.ReadAll(resp.Body)
    if (readErr != nil) {
        return Quote{}, readErr
    }

    return DecodeResponseJSON(body)

}