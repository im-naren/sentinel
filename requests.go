package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type Request struct {
	Timeout time.Duration
}

func (r *Request) buildGetURL(url string, params map[string]string) string {
	queryString := ""
	for k, v := range params {
		queryString += fmt.Sprintf("%s=%s&", k, v)
	}
	queryString = queryString[:len(queryString)-1]
	return fmt.Sprintf("%s?%s", url, queryString)
}


func (r *Request) get(url string, params map[string]string, headers map[string]string, cookies []*http.Cookie) (*http.Response, error) {
	client := &http.Client{}
	url = r.buildGetURL(url, params)

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

func (r *Request) fetchAndUnmarshal(url string, params map[string]string, headers map[string]string, cookies []*http.Cookie, result interface{}) error {
	response, err := r.get(url, params, headers, cookies)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	if strResult, ok := result.(*string); ok {
		*strResult = string(responseData)
		return nil
	}

	err = json.Unmarshal(responseData, result)
	if err != nil {
		return err
	}

	return nil
}

func (r *Request) buildPostPayload(payload map[string]string) string {
	payloadData := ""
	for k, v := range payload {
		payloadData += fmt.Sprintf("%s=%s&", k, v)
	}
	payloadData = payloadData[:len(payloadData)-1]
	return payloadData
}


func (r *Request) post(url string, payload map[string]string, headers map[string]string, cookies []*http.Cookie) (*http.Response, error) {
	payloadData := r.buildPostPayload(payload)

	req, err := http.NewRequest("POST", url, strings.NewReader(payloadData))
	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}

	client := &http.Client{}
	return client.Do(req)
}

func (r *Request) psotAndUnmarshal(url string, payload map[string]string, headers map[string]string, cookies []*http.Cookie, result interface{}) error {
	response, err := r.post(url, payload, headers, cookies)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	if strResult, ok := result.(*string); ok {
		*strResult = string(responseData)
		return nil
	}

	err = json.Unmarshal(responseData, result)
	if err != nil {
		return err
	}

	return nil
}	
