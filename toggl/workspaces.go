// ワークスペース
package toggl

import (
	"encoding/json"

	"github.com/mkohei/my-toggl/config"
)

const WORKSPACES_URL = "/api/v8/workspaces"

type Workspace struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func GetWorkspaces(conf config.Config) (workspaces []Workspace, err error) {
	url := BASE_URL + WORKSPACES_URL

	body, err := Get2(url, conf.APIToken)
	if err != nil {
		return workspaces, err
	}
	json.Unmarshal(body, &workspaces)
	return workspaces, nil
}
