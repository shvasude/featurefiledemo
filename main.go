package main

import (
	"fmt"
	"strings"

	lib "github.com/shvasude/featurefiledemo/lib"

	"gotest.tools/v3/icmd"
)

var (
	git = "github.com"
	//path = "*"
	path = "specs/pipelines"
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

	issueData := lib.IssueStruct{}

	updatedVal := "updated value"
	issueData.Fields.Description = &updatedVal
	lib.UpdateIssue("APPSVC-603", issueData)

	summary := "Title"
	key := "APPSVC"
	issueType := "Task"

	//issueData.Fields.Description = &gitURL
	issueData.Fields.Summary = &summary
	issueData.Fields.Project.Key = &key
	issueData.Fields.IssueType.Name = &issueType

	//lib.CreateIssue(issueData)

	//lib.GetIssue("APPSVC-603")

	//issueData.Reset()

	//updatedVal := "updated value"
	issueData.Fields.Description = &updatedVal
	lib.UpdateIssue("APPSVC-603", issueData)

	fmt.Printf("Execution Done")

	//This is only performed if you have an admin rights to delete a jira issue in a jira project
	//lib.DeleteIssue("APPSVC-603")

}
