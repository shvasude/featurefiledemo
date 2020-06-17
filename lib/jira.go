package lib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/BurntSushi/toml"
)

var (
	configFile  = "jira_config.toml"
	parameter   string
	apiEndPoint = "/rest/api/2/issue/"
	issue       IssueStruct
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

//Credentials a representation of a JIRA config which helds API permissions
type Credentials struct {
	Username string
	Password string
	URL      string
}

func (cred *Credentials) initConfig() {
	if _, err := os.Stat(configFile); err != nil {
		log.Fatalf("Error using config file: %v", err)
	}

	if _, err := toml.DecodeFile(configFile, cred); err != nil {
		log.Fatal("Error during decoding toml config: ", err)
	}
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

func sendRequest(jsonStr []byte, method string, url string) {
	cred := &Credentials{}
	cred.initConfig()
	fmt.Println("Json:", string(jsonStr))
	req, err := http.NewRequest(method, cred.URL+url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(cred.Username, cred.Password)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

}
