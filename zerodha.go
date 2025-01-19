package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

type Zerodha struct {
	userID   string
	password string
	twofa    string
	encToken string
	client   *http.Client
}

func NewZerodha(userID, password, twofa string) *Zerodha {
	jar, _ := cookiejar.New(nil)
	client := &http.Client{Jar: jar}
	return &Zerodha{
		userID:   userID,
		password: password,
		twofa:    twofa,
		client:   client,
	}
}

func (z *Zerodha) loginStep1() (map[string]interface{}, error) {
	resp, err := z.client.PostForm("https://kite.zerodha.com/api/login", url.Values{
		"user_id":  {z.userID},
		"password": {z.password},
	})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}

func (z *Zerodha) loginStep2(j map[string]interface{}) (map[string]interface{}, error) {
	data := url.Values{
		"user_id":     {z.userID},
		"request_id":  {j["data"].(map[string]interface{})["request_id"].(string)},
		"twofa_value": {z.twofa},
	}
	resp, err := z.client.PostForm("https://kite.zerodha.com/api/twofa", data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}

func (z *Zerodha) Login() (map[string]interface{}, error) {
	j, err := z.loginStep1()
	if err != nil {
		return nil, err
	}
	if j["status"] == "error" {
		return nil, errors.New(j["message"].(string))
	}

	j, err = z.loginStep2(j)
	if err != nil {
		return nil, err
	}
	if j["status"] == "error" {
		return nil, errors.New(j["message"].(string))
	}
	z.encToken = z.client.Jar.Cookies(nil)[0].Value
	return j, nil
}

// func main() {
//     z := NewZerodha("your_user_id", "your_password", "your_twofa")
//     result, err := z.Login()
//     if err != nil {
//         log.Fatal(err)
//     }
//     fmt.Println(result)
// }
