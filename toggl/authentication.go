// Authentication
package toggl

import (
	"encoding/json"

	"github.com/mkohei/my-toggl/config"
)

type AuthResult struct {
	Since int  `json:"since"`
	Data  Data `json:"data"`
}

type Data struct {
	ID         int         `json:"id"`
	APIToken   string      `json:"api_token"`
	Email      string      `json:"email"`
	FullName   string      `json:"fullname"`
	Workspaces []Workspace `json:"workspaces"`
}

const AUTH_URL = "/api/v8/me"

func GetMe(conf config.Config) (result AuthResult, err error) {
	url := BASE_URL + AUTH_URL

	body, err := Get2(url, conf.APIToken)
	if err != nil {
		return result, err
	}
	json.Unmarshal(body, &result)
	return result, nil
}
