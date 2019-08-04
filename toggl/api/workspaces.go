package api

import (
	"encoding/json"

	"github.com/mkohei/my-toggl/toggl"
)

// WorkspacesURL is endpoint to get workspace infomation
const WorkspacesURL = "/api/v8/workspaces"

// Workspace is one object in response array
type Workspace struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// GetWorkspaces request toggl Workspaces API
func GetWorkspaces(conf toggl.Config) (workspaces []Workspace, err error) {
	// Request
	body, err := toggl.GetUseBasicAuth(WorkspacesURL, conf.APIToken)
	if err != nil {
		return workspaces, err
	}
	json.Unmarshal(body, &workspaces)
	return workspaces, nil
}
