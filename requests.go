package main

import (
	"fmt"
	"net/http"
	"io"
	"log"
)


func get(url string) string {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
	}
	responseData, err := io.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }
	return string(responseData)
}

func main() {
	fmt.Println(get("https://www.google.com"))
}
