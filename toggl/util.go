package toggl

import (
	"io/ioutil"
	"net/http"
)

const BASE_URL = "https://www.toggl.com/"

func Get(url string) (body []byte, err error) {
	resp, err := http.Get(url)
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

func Get2(url string, apiToken string) (body []byte, err error) {
	req, _ := http.NewRequest("GET", url, nil)
	//seed := []byte(fmt.Sprintf("%s:api_token", apiToken))
	//sEnc := base64.StdEncoding.EncodeToString(seed)
	//req.Header.Set("Authorization", fmt.Sprintf("Basic %s", string(sEnc)))
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
