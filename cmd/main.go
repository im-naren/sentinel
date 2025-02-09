package main

import (
	"log"

	"github.com/im-naren/sentinel/config"
	integration "github.com/im-naren/sentinel/integrations"
)

func main() {
	// Set log flags
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	// Load configuration
	cfg, err := config.LoadConfig("config/dev.toml")
	if err != nil {
		log.Fatal(err)
	}

	nseLive := integration.NewNSELive()
	quote, err := nseLive.StockQuote(cfg.App.Name)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(quote)

	meta_info, err := nseLive.StockMeta(cfg.App.Name)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(meta_info)

	// z := NewZerodha()
	// result, err := z.Login()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println(result)
}
