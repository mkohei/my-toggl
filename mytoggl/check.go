package mytoggl

import (
	"fmt"
	"regexp"

	"github.com/mkohei/my-backlog-api-tool/backlog"

	"github.com/mkohei/my-backlog-api-tool/config"

	"github.com/mkohei/my-toggl/toggl/reports"
)

// CheckErrorCode is Enum for error code for check
type CheckErrorCode int

// content of CheckErrorCode
const (
	None CheckErrorCode = iota
	Project
	BacklogIssueWithProject
	ChildBacklogIssue
)

func (code CheckErrorCode) String() string {
	switch code {
	case Project:
		return "すべてのレコードにプロジェクトがあるか"
	case BacklogIssueWithProject:
		return "BacklogのIssueが指定されていること(設定したプロジェクトのみ)"
	case ChildBacklogIssue:
		return "子チケットが入力されていないこと"
	}
	return ""
}

// ShowErrorMessages display error message from code
func ShowErrorMessages(errorCodes []CheckErrorCode) {
	if len(errorCodes) == 0 {
		return
	}

	// Header
	fmt.Println("/// Error Messages")

	// 表示した回数を記憶(1度しか表示しない)
	codeCountMap := map[CheckErrorCode]int{}

	for _, code := range errorCodes {
		if codeCountMap[code] == 0 {
			fmt.Printf("[Code] %d | [Message] %s\n", code, code)
		}
		codeCountMap[code]++
	}
}

// CheckDetailedRecord check toggl Detailed Record with my rule
func CheckDetailedRecord(record reports.DetailedRecord) (ok bool, errorCode CheckErrorCode, err error) {
	// 設定の読み込み
	conf, err := LoadConfig()
	if err != nil {
		return false, None, err
	}

	// 対象外
	// * 入力忘れがないか
	// * 1日の稼働時間が正しいか（極端な時間がないか）

	// togglで完結
	// * すべてのレコードにプロジェクトがあるか
	ok = checkProject(record)
	if !ok {
		return false, Project, nil
	}

	// backlogも使用
	// * BacklogのIssueが必要なプロジェクトの確認
	ok, err = checkBacklogIssueWithProject(record, conf)
	if err != nil {
		return false, None, err
	}
	if !ok {
		return false, BacklogIssueWithProject, nil
	}
	// * 子チケットが入力されていないこと
	ok, err = checkChildBacklogIssue(record, conf)
	if err != nil {
		return false, None, err
	}
	if !ok {
		return false, ChildBacklogIssue, nil
	}

	return true, None, nil
}

// すべてのレコードにプロジェクトがあるか
func checkProject(record reports.DetailedRecord) bool {
	return record.Project != ""
}

// BacklogのIssueが指定されていること(設定したプロジェクトのみ)
func checkBacklogIssueWithProject(record reports.DetailedRecord, conf Config) (ok bool, err error) {
	if !isProjectNeededParentBacklogIssue(record, conf) {
		// チェック対象のプロジェクトではないのでOK
		return true, nil
	}
	// Backlog Issue Key を持つか確認
	issueKey := getBacklogIssueKey(record, conf)
	if issueKey == "" {
		return false, nil
	}
	return true, nil
}

func isProjectNeededParentBacklogIssue(record reports.DetailedRecord, conf Config) bool {
	// 対象のプロジェクト
	targetProjects := conf.ToggleProjectsRequiredBacklogIssue

	for _, project := range targetProjects {
		if record.Project == project {
			return true
		}
	}
	return false
}

// 子チケットが入力されていないこと
func checkChildBacklogIssue(record reports.DetailedRecord, conf Config) (hasNotChildBacklogIssue bool, err error) {
	backlogIssueKey := getBacklogIssueKey(record, conf)
	if backlogIssueKey == "" {
		// 課題キーを持たないので true
		return true, nil
	}
	return isParentBacklogIssue(backlogIssueKey)
}

func hasParentBacklogIssueKey(record reports.DetailedRecord, conf Config) (hasParentBacklogIssueKey bool, err error) {
	backlogIssueKey := getBacklogIssueKey(record, conf)
	if backlogIssueKey == "" {
		// 課題キーを持たないので false
		return false, nil
	}
	return isParentBacklogIssue(backlogIssueKey)
}

func getBacklogIssueKey(record reports.DetailedRecord, conf Config) (backlogIssueKey string) {
	// 対象のBacklog プロジェクトキー
	backlogProjectKeys := conf.BacklogProjectKeys

	// 正規表現作成
	regEx := ``
	for i, projectKey := range backlogProjectKeys {
		if i == 0 {
			regEx = `(` + projectKey + `)`
		} else {
			regEx = regEx + `|(` + projectKey + `)`
		}
	}
	regEx = `(` + regEx + `)-\d+`

	r := regexp.MustCompile(regEx)
	return r.FindString(record.Description)
}

func isParentBacklogIssue(backlogIssueKey string) (isParent bool, err error) {
	// Backlog用設定
	conf, err := config.LoadConfig()
	if err != nil {
		return isParent, err
	}

	// 課題情報の取得
	issue, err := backlog.GetIssue(conf, backlogIssueKey)
	if err != nil {
		return isParent, err
	}

	// 親課題かどうか
	return issue.ParentIssueID == 0, nil
}
