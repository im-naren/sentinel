// nse.go
package main

import (
	"encoding/json"
	"log"
	"time"
)

type NSELive struct {
	client  *Request
	baseURL string
	pageURL string
	routes  map[string]string
}

func NewNSELive() *NSELive {
	client := &Request{
		Timeout: 5 * time.Second,
	}

	return &NSELive{
		client:  client,
		baseURL: "https://www.nseindia.com/api",
		pageURL: "https://www.nseindia.com/get-quotes/equity?symbol=LT",
		routes: map[string]string{
			"stock_meta":                 "/equity-meta-info",
			"stock_quote":                "/quote-equity",
			"derivative_quote":           "/quote-derivative",
			"market_status":              "/marketStatus",
			"chart_data":                 "/chart-databyindex",
			"market_turnover":            "/market-turnover",
			"equity_derivative_turnover": "/equity-stock",
			"all_indices":                "/allIndices",
			"live_index":                 "/equity-stockIndices",
			"index_option_chain":         "/option-chain-indices",
			"equity_option_chain":        "/option-chain-equities",
			"currency_option_chain":      "/option-chain-currency",
			"pre_open_market":            "/market-data-pre-open",
			"holiday_list":               "/holiday-master?type=trading",
			"corporate_announcements":    "/corporate-announcements",
		},
	}
}

func (n *NSELive) stock_quote(symbol string) (map[string]interface{}, error) {
	payload := map[string]string{
		"symbol": symbol,
	}
	return n.get("stock_quote", payload)
}

func (n *NSELive) get(route string, payload map[string]string) (map[string]interface{}, error) {
	url := n.baseURL + n.routes[route]

	var result map[string]interface{}
	err := n.client.fetchAndUnmarshal(url, payload, nil, nil, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}