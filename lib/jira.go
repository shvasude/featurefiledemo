package lib

import (	
	"fmt"
	"encoding/json"
	"log"
)

//Project defines the project values
type Project struct {
	Key *string `json:"key,omitempty"`
}

//IssueType defines the issue type values
type IssueType struct {
	Name *string `json:"name,omitempty"`
}

//Priority defines the priority values
type Priority struct {
	ID *string `json:"id,omitempty"`
}

//IssueStruct is a representation of a Jira Issue
type IssueStruct struct {
	Fields struct {
		Project     *Project   `json:"project,omitempty"`
		Summary     *string    `json:"summary,omitempty"`
		Description *string    `json:"description,omitempty"`
		IssueType   *IssueType `json:"issuetype,omitempty"`
		Priority    *Priority  `json:"priority,omitempty"`
	} `json:"fields,omitempty"`
}

//CreateIssue creates a new issue
func CreateIssue(issueData IssueStruct) {
	fmt.Println("Creating new issue.")
	marshalledIssue, err := json.Marshal(issueData)
	if err != nil {
		log.Fatal("Error occured when Marshaling Issue:", err)
	}
	sendRequest(marshalledIssue, "POST", apiEndPoint)
}

//UpdateIssue update an existing issue; PUT /rest/api/2/issue/{issueIdOrKey}
func UpdateIssue(issueID string, issueData IssueStruct) {
	fmt.Println("Updating existing issue.")
	apiURL := fmt.Sprintf("%s%s", apiEndPoint, issueID)
	marshalledIssue, err := json.Marshal(issueData)
	if err != nil {
		log.Fatal("Error occured when Marshaling Issue:", err)
	}
	sendRequest(marshalledIssue, "PUT", apiURL)
}

//GetIssue gets the information of a jira issue; GET /rest/api/2/issue/{issueIdOrKey}
func GetIssue(issueID string) {
	apiURL := fmt.Sprintf("%s%s", apiEndPoint, issueID)
	sendRequest(nil, "GET", apiURL)
}

//DeleteIssue deletes the issue in jira; DELETE /rest/api/2/issue/{issueIdOrKey}
func DeleteIssue(issueID string) {
	apiURL := fmt.Sprintf("%s%s", apiEndPoint, issueID)
	sendRequest(nil, "DELETE", apiURL)
}