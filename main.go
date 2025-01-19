package main

import (
	"log"
)

func main() {

	// // Set log flags
	// log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	// nseLive := integration.NewNSELive()
	// quote, err := nseLive.StockQuote("LT")

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Println(quote)

	// meta_info, err := nseLive.StockMeta("LT")

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Println(meta_info)

	z := NewZerodha()
	result, err := z.Login()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(result)

}
