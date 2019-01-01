package toggl

import (
	"encoding/json"
	"io/ioutil"
)

// Config is for setting
type Config struct {
	APIToken      string `json:"api_token"`
	Email         string `json:"email"`
	WorkspaceName string `json:"workspace_name"`
	WorkspaceID   int    `json:"workspace_id"` // TODO: for test
	UserID        int    `json:"user_id"`      // TODO: for test
}

// LoadConfig load conf.json and return Config
func LoadConfig() (conf Config, err error) {
	raw, err := ioutil.ReadFile("./conf.json")
	if err != nil {
		return conf, err
	}
	json.Unmarshal(raw, &conf)
	return conf, nil
}
