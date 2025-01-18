package main

type MarketMaven interface {
	stock_quote(symbol string) (map[string]interface{}, error)
	// stock_meta(symbol string) (map[string]interface{}, error)
}
