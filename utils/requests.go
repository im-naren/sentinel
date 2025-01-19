package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"strings"
	"time"
)

type Request struct {
	Client  *http.Client
	Timeout time.Duration
}

func NewRequest(timeout time.Duration) *Request {
	jar, _ := cookiejar.New(nil) // Create a cookie jar to maintain session
	client := &http.Client{
		Timeout: timeout,
		Jar:     jar,
	}
	return &Request{
		Client:  client,
		Timeout: timeout,
	}
}

func (r *Request) buildGetURL(url string, params map[string]string) string {
	if params == nil {
		return url
	}
	
	queryString := ""
	for k, v := range params {
		queryString += fmt.Sprintf("%s=%s&", k, v)
	}
	queryString = queryString[:len(queryString)-1]
	return fmt.Sprintf("%s?%s", url, queryString)
}

func (r *Request) Get(url string, params map[string]string, headers map[string]string, cookies []*http.Cookie) (*http.Response, error) {
	fullURL := r.buildGetURL(url, params)
	log.Println(fullURL)

	req, err := http.NewRequest("GET", fullURL, nil)
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

	return r.Client.Do(req)
}

func (r *Request) FetchAndUnmarshal(url string, params map[string]string, headers map[string]string, cookies []*http.Cookie, result interface{}) error {
	response, err := r.Get(url, params, headers, cookies)
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

	return r.Client.Do(req)
}

func (r *Request) PostAndUnmarshal(url string, payload map[string]string, headers map[string]string, cookies []*http.Cookie, result interface{}) error {
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
