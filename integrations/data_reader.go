package integration

type DataReader interface {
	StockQuote(symbol string) (map[string]interface{}, error)
	// stock_meta(symbol string) (map[string]interface{}, error)
}
