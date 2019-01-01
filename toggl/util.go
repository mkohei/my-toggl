package toggl

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// BaseURL is base endopoint
const BaseURL = "https://www.toggl.com"

// Get request http
func Get(url string) (body []byte, err error) {
	resp, err := http.Get(BaseURL + url)
	if err != nil {
		return body, err
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return body, err
	}
	return body, nil
}

// GetUseBasicAuth request http with basic auth
func GetUseBasicAuth(url string, apiToken string) (body []byte, err error) {
	// TODO:
	fmt.Println(url)

	req, _ := http.NewRequest("GET", BaseURL+url, nil)
	req.SetBasicAuth(apiToken, "api_token")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return body, err
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return body, err
	}
	return body, nil
}
