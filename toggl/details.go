// reports-details
package toggl

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/mkohei/my-toggl/config"
)

type DetailsResult struct {
	TotalGrand int `json:"total_grand"`
	TotalCount int `json:"total_count"`
	PerPage    int `json:"per_page"`
}

const DETAILS_URL = "/reports/api/v2/details"

func GetDetails(conf config.Config) (result DetailsResult, err error) {
	values := url.Values{}
	values.Add("user_agent", conf.Email)
	values.Add("workspace_id", strconv.Itoa(conf.WorkspaceID))
	values.Add("user_ids[]", strconv.Itoa(conf.UserID))

	url := BASE_URL + DETAILS_URL + "?" + values.Encode()
	body, err := Get2(url, conf.APIToken)
	if err != nil {
		return result, err
	}

	json.Unmarshal(body, &result)
	return result, nil
}
