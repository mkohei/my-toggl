package main

import (
	"fmt"
	"strconv"

	"github.com/mkohei/my-toggl/toggl"

	"github.com/mkohei/my-toggl/config"
)

func errorExit(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func main() {
	conf, err := config.LoadConfig()
	errorExit(err)

	// workspaces
	/*
		workspaces, err := toggl.GetWorkspaces(conf)
		errorExit(err)
		fmt.Println(workspaces)
	*/

	// me (auth)
	/*
		result, err := toggl.GetMe(conf)
		errorExit(err)
		fmt.Printf("%+v\n", result)
	*/

	// reports/details
	/*
		URL := "https://www.toggl.com/reports/api/v2/details" + "?" + conf.Param
		body, err := toggl.Get2(URL, conf.APIToken)
		errorExit(err)
		fmt.Println(string(body))
	*/

	params := map[string]string{}
	page := 0
	for {
		params["page"] = strconv.Itoa(page)
		result, err := toggl.GetDetailsMonth(conf, "2019-02", params)
		errorExit(err)

		// 処理
		fmt.Println(len(result.Data), result.PerPage, result.TotalCount, result.TotalCount/result.PerPage)
		// 処理

		page++
		if page >= result.TotalCount/result.PerPage {
			break
		}
	}
}
