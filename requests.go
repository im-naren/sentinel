package main

import (
	"encoding/json"
	"io"
	"net/http"
)

func get(url string, payload map[string]string, headers map[string]string, cookies []*http.Cookie) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Add headers if any
	for key, value := range headers {
		req.Header.Add(key, value)
	}

	// Add cookies if any
	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}

	return client.Do(req)
}

func FetchAndUnmarshal[T any](url string, payload map[string]string, headers map[string]string, cookies []*http.Cookie) (*T, error) {
	response, err := get(url, payload, headers, cookies)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if _, ok := any(*new(T)).(string); ok {
		result := any(string(responseData)).(T)
		return &result, nil
	}

	var result T
	err = json.Unmarshal(responseData, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
