package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/shvasude/featurefiledemo/lib"

	"gotest.tools/v3/icmd"
)

var (
	git  = "github.com"
	path = "specs/pipelines"
)

var flags struct {
	Comment     string
	Description string
	IssueKey    string
	Priority    string
	Title       string
	Project     string
	gitURL      string
}

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

func init() {

	flag.StringVar(&flags.Comment, "m", "Default Comment", "A Comment when changing the status of an Issue.")
	flag.StringVar(&flags.Description, "d", "Default Description", "Provide a description for a created Issue.")
	flag.StringVar(&flags.Priority, "p", "2", "The priority of an Issue which will be set.")
	flag.StringVar(&flags.IssueKey, "k", "", "Issue key/ID of an issue.")
	flag.StringVar(&flags.Title, "t", "Default Title", "Title of an Issue.")
	flag.StringVar(&flags.Project, "o", "IT", "Define a Project to create a ticket in.")
	flag.Parse()
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

	if len(flag.Args()) < 1 {
		log.Fatal("Please provide an action to take. Usage information:")
	}
	parameter := flag.Arg(0)
	switch parameter {
	case "close":
		lib.DeleteIssue(flags.IssueKey)
	case "start":
		lib.UpdateIssue(flags.IssueKey, flags.gitURL)
	case "create":
		lib.CreateIssue(flags.gitURL)
	}

}
