package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	APIToken string `json:"api_token"`
	//Param    string `json:"param"`
	Email       string `json:"email"`
	UserID      int    `json:"user_id"`
	WorkspaceID int    `json:"workspace_id"`
}

func LoadConfig() (conf Config, err error) {
	raw, err := ioutil.ReadFile("./conf.json")
	if err != nil {
		return conf, err
	}
	json.Unmarshal(raw, &conf)
	return conf, nil
}
