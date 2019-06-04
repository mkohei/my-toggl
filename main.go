package main

import (
	"fmt"

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
	result, err := toggl.GetDetails(conf)
	errorExit(err)
	fmt.Println(result)

}
