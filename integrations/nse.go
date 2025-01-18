// nse.go
package integration

import (
	"time"

	"github.com/im-naren/sentinel/utils"
)

type NSELive struct {
	client  *utils.Request
	baseURL string
	pageURL string
	routes  map[string]string
	headers map[string]string
}

func initialize_nse(client *utils.Request, headers map[string]string) {
	url := "https://www.nseindia.com"
	client.Get(url, nil, headers, nil)
}

func NewNSELive() *NSELive {
	client := utils.NewRequest(5 * time.Second)
	headers := map[string]string{
		"accept":             "*/*",
		"accept-language":    "en-IN,en;q=0.9,hi-IN;q=0.8,hi;q=0.7,en-GB;q=0.6,en-US;q=0.5",
		"priority":           "u=1, i",
		"referer":            "https://www.nseindia.com/market-data/live-equity-market?symbol=NIFTY%2050",
		"sec-ch-ua":          `"Chromium";v="124", "Google Chrome";v="124", "Not-A.Brand";v="99"`,
		"sec-ch-ua-mobile":   "?0",
		"sec-ch-ua-platform": `"macOS"`,
		"sec-fetch-dest":     "empty",
		"sec-fetch-mode":     "cors",
		"sec-fetch-site":     "same-origin",
		"user-agent":         "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36",
	}
	initialize_nse(client, headers)

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
		headers: headers,
	}
}

func (n *NSELive) StockQuote(symbol string) (map[string]interface{}, error) {
	payload := map[string]string{
		"symbol": symbol,
	}
	return n.get("stock_quote", payload)
}

func (n *NSELive) get(route string, payload map[string]string) (map[string]interface{}, error) {
	url := n.baseURL + n.routes[route]

	var result map[string]interface{}
	err := n.client.FetchAndUnmarshal(url, payload, n.headers, nil, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
