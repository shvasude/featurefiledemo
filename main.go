package main

import (
	"fmt"
	"strings"

	lib "github.com/shvasude/featurefiledemo/lib"
	"gotest.tools/v3/icmd"
)

var (
	git  = "github.com"
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
	issueData.Fields.Description = &gitURL
	lib.UpdateIssue("APPSVC-608", issueData)

	summary := "NewTest Title"
	key := "APPSVC"
	Project := lib.Project{Key: &key}
	issueType := "Task"
	Issuetype := lib.IssueType{Name: &issueType}
	priority := "10004"
	Priority := lib.Priority{ID: &priority}

	issueData.Fields.Description = &gitURL
	issueData.Fields.Summary = &summary
	issueData.Fields.Project = &Project
	issueData.Fields.IssueType = &Issuetype
	issueData.Fields.Priority = &Priority

	lib.CreateIssue(issueData)

	fmt.Printf("Execution Done")

}
