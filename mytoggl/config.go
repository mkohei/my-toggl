package mytoggl

import (
	"encoding/json"
	"io/ioutil"
)

// Config is settings for toggl check
type Config struct {
	BacklogProjectKeys                 []string `json:"backlog_project_keys"`
	ToggleProjectsRequiredBacklogIssue []string `json:"toggl_projects_required_backlog_issue"`
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
