package main

import (
	"log"

	integration "github.com/im-naren/sentinel/integrations"
)

func main() {

	// Set log flags
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	nseLive := integration.NewNSELive()
	quote, err := nseLive.StockQuote("LT")

	if err != nil {
		log.Fatal(err)
	}

	log.Println(quote)

}
