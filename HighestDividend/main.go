package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type Divident struct {
	Ticker string  `json:"ticker"`
	Divid  float64 `json:"dividend"`
}

func main() {
	fmt.Println("Stock Price AI")

	stocks_json := `[
    {"ticker":"APPL","dividend":0.5},
    {"ticker":"GOOG","dividend":0.2},
    {"ticker":"MSFT","dividend":0.3}
  ]`

	highestDividend:= HighestDividend(stocks_json)
	fmt.Println(highestDividend)
	// dividents := []Dividents{
	// 	{Ticker: "APPL", Dividend: 0.5},
	// 	{Ticker: "MSFT", Dividend: 0.2},
	//}
}
func HighestDividend(stocks_json string) string {

	d := []Divident{}
	err := json.NewDecoder(bytes.NewBuffer([]byte(stocks_json))).Decode(&d)
	if err != nil {
		log.Fatal(err.Error())
		return ""
	}
	max, ticker := 0.0, ""
	if len(d) > 0 {
		for _, v := range d {
			if v.Divid > max {
				max = v.Divid
				ticker = v.Ticker
			}
		}
	}
	return ticker
}
