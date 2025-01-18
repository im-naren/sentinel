package main

type Maven interface {
	stock_quote(symbol string) (map[string]interface{}, error)
	// stock_meta(symbol string) (map[string]interface{}, error)
}
