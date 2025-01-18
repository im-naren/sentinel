package main

import (
	"log"
)

func main() {
	// url := "https://www.nseindia.com/json/gainerLossersValue.json"
	// headers := map[string]string{
	// 	"accept":             "*/*",
	// 	"accept-language":    "en-IN,en;q=0.9,hi-IN;q=0.8,hi;q=0.7,en-GB;q=0.6,en-US;q=0.5",
	// 	"priority":           "u=1, i",
	// 	"referer":            "https://www.nseindia.com/market-data/live-equity-market?symbol=NIFTY%2050",
	// 	"sec-ch-ua":          "\"Chromium\";v=\"124\", \"Google Chrome\";v=\"124\", \"Not-A.Brand\";v=\"99\"",
	// 	"sec-ch-ua-mobile":   "?0",
	// 	"sec-ch-ua-platform": "\"macOS\"",
	// 	"sec-fetch-dest":     "empty",
	// 	"sec-fetch-mode":     "cors",
	// 	"sec-fetch-site":     "same-origin",
	// 	"user-agent":         "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36",
	// }
	// resp, err := FetchAndUnmarshal[string](url, nil, headers, nil)
	// if err != nil {
	// 	fmt.Print(err.Error())
	// 	os.Exit(1)
	// }
	// fmt.Println(*resp)

	nseLive := NewNSELive()
	quote, err := nseLive.stock_quote("LT")

	if err != nil {
		log.Fatal(err)
	}

	log.Println(quote)

}
