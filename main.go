package main

import (
	"log"
)

func main() {

	// Set log flags
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	nseLive := NewNSELive()
	quote, err := nseLive.stock_quote("LT")

	if err != nil {
		log.Fatal(err)
	}

	log.Println(quote)

}
