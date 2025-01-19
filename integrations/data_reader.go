package integration

type DataReader interface {
	StockQuote(symbol string) (map[string]interface{}, error)
	StockMeta(symbol string) (map[string]interface{}, error)
}
