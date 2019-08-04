package api

import (
	"encoding/json"

	"github.com/mkohei/my-toggl/toggl"
)

// AuthURL is endpoint for getting my information
const AuthURL = "/api/v8/me"

// AuthResponse is response to
type AuthResponse struct {
	Since int  `json:"since"`
	Data  Data `json:"data"`
}

// Data is content in AuthResponse
type Data struct {
	ID         int         `json:"id"`
	APIToken   string      `json:"api_token"`
	Email      string      `json:"email"`
	FullName   string      `json:"fullname"`
	Workspaces []Workspace `json:"workspaces"`
}

// GetAboutMeWithAPIToken request toggl Authentication API
func GetAboutMeWithAPIToken(conf toggl.Config) (response AuthResponse, err error) {
	// Request
	body, err := toggl.GetUseBasicAuth(AuthURL, conf.APIToken)

	if err != nil {
		return response, err
	}

	json.Unmarshal(body, &response)
	return response, nil
}
