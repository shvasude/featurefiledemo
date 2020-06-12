package main

import (
	"fmt"
	"strings"

	"github.com/shvasude/featurefiledemo/lib"

	"gotest.tools/v3/icmd"
)

var (
	git = "github.com"
	//path = "*"
	path = "feature"
)

//Run function executes a command with timeout
func run(cmd ...string) *icmd.Result {
	currentCmd := icmd.Cmd{
		Command: cmd,
	}
	fmt.Printf("=> Command to execute: %v \n", currentCmd)
	return icmd.RunCmd(currentCmd)
}

func getBranch() string {
	return strings.TrimSpace(run("git", "rev-parse", "--abbrev-ref", "HEAD").Stdout())
}

func main() {

	res := strings.TrimSpace(run("pwd").Stdout())
	lstArr := strings.Split(res, git)

	gitURL := fmt.Sprintf("http://%s/%s/tree/%s/%s", git, lstArr[1], getBranch(), path)
	fmt.Printf("Git URL - %s", gitURL)
	//lib.CreateIssue(gitURL)
	lib.GetIssue("APPSVC-603")
	lib.UpdateIssue("APPSVC-603", gitURL)

	fmt.Printf("Execution Done")

	//This is only performed if you have an admin rights to delete a jira issue in a jira project
	//lib.DeleteIssue("APPSVC-603")

}
