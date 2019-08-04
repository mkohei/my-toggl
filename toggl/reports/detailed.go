package reports

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"github.com/jinzhu/now"

	"github.com/mkohei/my-toggl/toggl"
)

// DetailedURL is endpoinst to get Detailed Report
const DetailedURL = "/reports/api/v2/details"

// DetailedResponse is response
type DetailedResponse struct {
	TotalGrand int              `json:"total_grand"`
	TotalCount int              `json:"total_count"`
	PerPage    int              `json:"per_page"`
	Data       []DetailedRecord `json:"data"`
}

// DetailedRecord is a record in DetailedResponse
type DetailedRecord struct {
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

// GetDetailed request toggl Detailed Report API
func GetDetailed(conf toggl.Config, params map[string]string) (response DetailedResponse, err error) {
	// パラメタ準備
	values := url.Values{}
	values.Add("user_agent", conf.Email)
	for key, val := range params {
		values.Add(key, val)
	}

	// URL作成
	url := DetailedURL + "?" + values.Encode()

	// Request
	body, err := toggl.GetUseBasicAuth(url, conf.APIToken)

	if err != nil {
		return response, err
	}

	json.Unmarshal(body, &response)
	return response, nil
}

// GetDetailedAll return all DetailedRecords
func GetDetailedAll(conf toggl.Config, params map[string]string) (detailedRecords []DetailedRecord, err error) {

	page := 1
	for {
		params["page"] = strconv.Itoa(page)

		// Request
		detailedResponse, err := GetDetailed(conf, params)
		if err != nil {
			return detailedRecords, err
		}

		// 取得したレコードが空なら終了
		if len(detailedResponse.Data) == 0 {
			break
		}

		// 更新
		detailedRecords = append(detailedRecords, detailedResponse.Data...)
		page++
	}
	return detailedRecords, nil
}

// GetDetailedMonth return Detailed Reports in any month
func GetDetailedMonth(conf toggl.Config, targetMonth string, params map[string]string) (detailedRecords []DetailedRecord, err error) {
	// 日付リクエストパラメタ
	t, err := now.Parse(targetMonth)
	if err != nil {
		return detailedRecords, err
	}
	params["since"] = fmt.Sprintf("%v", now.New(t).BeginningOfMonth().Format("2006-01-02"))
	params["until"] = fmt.Sprintf("%v", now.New(t).EndOfMonth().Format("2006-01-02"))

	// Request
	detailedRecords, err = GetDetailedAll(conf, params)
	if err != nil {
		return detailedRecords, err
	}
	return detailedRecords, nil
}

// type Records []Record

// func (records Records) Len() int {
// 	return len(records)
// }
// func (records Records) Swap(i, j int) {
// 	records[i], records[j] = records[j], records[i]
// }
// func (records Records) Less(i, j int) bool {
// 	return records[i].ID < records[j].ID
// }

// func GetDetails(conf config.Config, params map[string]string) (result DetailsResult, err error) {
// 	values := url.Values{}
// 	values.Add("user_agent", conf.Email)
// 	for key, val := range params {
// 		values.Add(key, val)
// 	}

// 	url := BASE_URL + DETAILS_URL + "?" + values.Encode()

// 	body, err := Get2(url, conf.APIToken)
// 	if err != nil {
// 		return result, err
// 	}

// 	json.Unmarshal(body, &result)
// 	return result, nil
// }

// // GetDetailsMonth provides to get details in a month with paging
// func GetDetailsMonth(conf config.Config, targetMonth string, params map[string]string) (result DetailsResult, err error) {
// 	t, err := now.Parse(targetMonth)
// 	if err != nil {
// 		return result, err
// 	}
// 	params["since"] = fmt.Sprintf("%v", now.New(t).BeginningOfMonth().Format("2006-01-02"))
// 	params["until"] = fmt.Sprintf("%v", now.New(t).EndOfMonth().Format("2006-01-02"))

// 	return GetDetails(conf, params)
// }

// // GetDetailsMonthAll provides to get all details in a month
// func GetDetailsMonthAll(conf config.Config, targetMonth string, params map[string]string) (records []Record, err error) {
// 	t, err := now.Parse(targetMonth)
// 	if err != nil {
// 		return records, err
// 	}
// 	params["since"] = fmt.Sprintf("%v", now.New(t).BeginningOfMonth().Format("2006-01-02"))
// 	params["until"] = fmt.Sprintf("%v", now.New(t).EndOfMonth().Format("2006-01-02"))

// 	page := 1
// 	idmap := map[int]Record{}
// 	for {
// 		params["page"] = strconv.Itoa(page)
// 		result, err := GetDetailsMonth(conf, targetMonth, params)
// 		if err != nil {
// 			return records, err
// 		}
// 		records = append(records, result.Data...)
// 		for _, record := range result.Data {
// 			idmap[record.ID] = record
// 		}

// 		page++
// 		// 最後のページに行けば終わり
// 		if float64(page)-1 >= float64(result.TotalCount)/float64(result.PerPage) {
// 			break
// 		}
// 	}
// 	return records, nil
// }
