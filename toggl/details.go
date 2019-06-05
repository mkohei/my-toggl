// reports-details
package toggl

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"github.com/jinzhu/now"

	"github.com/mkohei/my-toggl/config"
)

type DetailsResult struct {
	TotalGrand int      `json:"total_grand"`
	TotalCount int      `json:"total_count"`
	PerPage    int      `json:"per_page"`
	Data       []Record `json:"data"`
}

type Record struct {
	ID          int      `json:"id"`
	PID         int      `json:"pid"`
	TID         int      `json:"tid"`
	UID         int      `json:"uid"`
	Description string   `json:"description"`
	Start       string   `json:"start"`
	End         string   `json:"end"`
	User        string   `json:"user"`
	Project     string   `json:"project"`
	Tags        []string `json:"tags"`
}

const DETAILS_URL = "/reports/api/v2/details"

func GetDetails(conf config.Config, params map[string]string) (result DetailsResult, err error) {
	values := url.Values{}
	values.Add("user_agent", conf.Email)
	values.Add("workspace_id", strconv.Itoa(conf.WorkspaceID))
	values.Add("user_ids", strconv.Itoa(conf.UserID))
	for key, val := range params {
		values.Add(key, val)
	}

	url := BASE_URL + DETAILS_URL + "?" + values.Encode()
	fmt.Println(url)
	body, err := Get2(url, conf.APIToken)
	if err != nil {
		return result, err
	}

	json.Unmarshal(body, &result)
	return result, nil
}

func GetDetailsMonth(conf config.Config, targetMonth string, params map[string]string) (result DetailsResult, err error) {
	t, err := now.Parse(targetMonth)
	if err != nil {
		return result, err
	}
	params["since"] = fmt.Sprintf("%v", now.New(t).BeginningOfMonth().Format("2006-01-02"))
	params["util"] = fmt.Sprintf("%v", now.New(t).EndOfMonth().Format("2006-01-02"))

	fmt.Println(params)

	return GetDetails(conf, params)
}
