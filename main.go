package main

import (
	"fmt"
	"strconv"

	"github.com/mkohei/my-toggl/mytoggl"
	"github.com/mkohei/my-toggl/toggl"
	"github.com/mkohei/my-toggl/toggl/api"
	"github.com/mkohei/my-toggl/toggl/reports"
)

func main() {
	conf, err := toggl.LoadConfig()
	errorExit(err)

	targetMonth := "2019-08"

	// Prepare Request
	ids, err := getNeededTogglIDs(conf)
	errorExit(err)
	requestParams := map[string]string{}
	requestParams["workspace_id"] = strconv.Itoa(ids["workspace_id"])
	requestParams["user_ids"] = strconv.Itoa(ids["user_id"])
	// requestParams["workspace_id"] = strconv.Itoa(conf.UserID)
	// requestParams["user_ids"] = strconv.Itoa(conf.WorkspaceID)

	// Togglデータの取得
	detailedRecords, err := reports.GetDetailedMonth(conf, targetMonth, requestParams)
	errorExit(err)

	// チェック
	ngCount := 0
	errorCodes := []mytoggl.CheckErrorCode{}
	for _, record := range detailedRecords {
		ok, errorCode, err := mytoggl.CheckDetailedRecord(record)
		if err != nil {
			errorExit(err)
		}
		if ok {
			// 特に何もしなくていいかな...
		} else {
			showCheckErrorRecord(record, errorCode)
			ngCount++
			errorCodes = append(errorCodes, errorCode)
		}
	}
	// 全体結果の表示
	mytoggl.ShowErrorMessages(errorCodes)
	fmt.Printf("*****\nNG : %v / %v\n", ngCount, len(detailedRecords))
}

// Request Toggl AUthentication API to get user_id, workspace_id
func getNeededTogglIDs(conf toggl.Config) (ids map[string]int, err error) {
	ids = map[string]int{}

	// conf が持っている場合はそれを使用
	if conf.UserID != 0 && conf.WorkspaceID != 0 {
		ids["user_id"] = conf.UserID
		ids["workspace_id"] = conf.WorkspaceID
	}

	// Request
	authResponse, err := api.GetAboutMeWithAPIToken(conf)
	if err != nil {
		return ids, err
	}
	ids["user_id"] = authResponse.Data.ID

	// workspace_name から対象のworkspaceを探す
	for _, workspace := range authResponse.Data.Workspaces {
		if workspace.Name == conf.WorkspaceName {
			ids["workspace_id"] = workspace.ID
			break
		}
	}
	return ids, nil
}

func showCheckErrorRecord(record reports.DetailedRecord, errorCode mytoggl.CheckErrorCode) {
	fmt.Printf("[ErrorCode] %d | [Description] %-33s\n", errorCode, record.Description)
}

// errorExit is to error handling in main
func errorExit(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
