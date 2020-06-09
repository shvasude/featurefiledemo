package main

import (
	"fmt"
	"strings"

	"gotest.tools/v3/icmd"
)

var (
	git = "github.com"
	//path = "*"
	path = "specs/pipelines"
)

//https://github.com/<<project>>/<<repo>>/tree/<<branch>>/<<path>>

//Run function executes a command with timeout
func run(cmd ...string) *icmd.Result {
	currentCmd := icmd.Cmd{
		Command: cmd,
	}
	fmt.Printf("=> Command to execute: %v \n", currentCmd)
	return icmd.RunCmd(currentCmd)
}

func getBranch() string {
	//git rev-parse --abbrev-ref HEAD
	branch := strings.TrimSpace(run("git", "rev-parse", "--abbrev-ref", "HEAD").Stdout())
	//url := fmt.Sprintf("%s", branch)
	return branch
}

func main() {

	res := strings.TrimSpace(run("pwd").Stdout())
	lstArr := strings.Split(res, git)

	url := fmt.Sprintf("https://%s/%s", git, lstArr[1])
	fmt.Printf(url)
	fmt.Printf(getBranch())
	url1 := fmt.Sprintf("%s/tree/%s", url, getBranch())
	url2 := fmt.Sprintf("%s/%s", url1, path)
	fmt.Printf(url2)
	//examplePath := fmt.Sprintf("%s/%s", getExamplesDir())

	//appStatusEndpoint := fmt.Sprintf("http://%s/api/status/dbNameCM")
	///checkNodeJSAppFrontend(appStatusEndpoint, dbName)
}
