package main

import (
	"fmt"
	"net/http"
	"io"
	"log"
)

func buildGetURL(url string, params map[string]string) string {
    queryString := ""
    for k, v := range params {
        queryString += fmt.Sprintf("%s=%s&", k, v)
    }
    queryString = queryString[:len(queryString)-1]
    return fmt.Sprintf("%s?%s", url, queryString)
}

func get(url string, params map[string]string) string {
    _url := buildGetURL(url, params)
    response, err := http.Get(_url)
    if err != nil {
        fmt.Println("Error:", err)
        return ""
    }
    defer response.Body.Close()
    responseData, err := io.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }
    return string(responseData)
}

func main() {
    fmt.Println(get("https://www.google.com", map[string]string{}))
}
