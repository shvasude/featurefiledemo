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
	issue       Issue
	apiEndPoint = "/rest/api/2/issue/"
)

//Issue is a representation of a Jira Issue
type Issue struct {
	Fields struct {
		Project struct {
			Key string `json:"key,omitempty"`
		} `json:"project,omitempty"`
		Summary     string `json:"summary,omitempty"`
		Description string `json:"description,omitempty"`
		Issuetype   struct {
			Name string `json:"name,omitempty"`
		} `json:"issuetype,omitempty"`
		Priority struct {
			ID string `json:"priority,omitempty"`
		} `json:"priority,omitempty"`
	} `json:"omitempty,omitempty"`
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
func CreateIssue(gitURL string) {
	fmt.Println("Creating new issue.")
	issue.Fields.Summary = "Title"
	issue.Fields.Project.Key = "APPSVC"
	issue.Fields.Issuetype.Name = "Task"
	issue.Fields.Description = gitURL
	issue.Fields.Priority.ID = "4"
	marshalledIssue, err := json.Marshal(issue)
	if err != nil {
		log.Fatal("Error occured when Marshaling Issue:", err)
	}
	sendRequest(marshalledIssue, "POST", "")
}

/*Need to ignore the unwanted struct for marshalling as
From post man, gives 200 on {"fields":{"description":"https://github.com//shvasude/featurefiledemo/tree/master/feature"}}
From post man, gives gives 400 error that says 'could not find issuetype by id or name
*/

//UpdateIssue update an existing issue; PUT /rest/api/2/issue/{issueIdOrKey}
func UpdateIssue(issueID string, gitURL string) {
	fmt.Println("Updating existing issue.")
	apiURL := fmt.Sprintf("%s%s", apiEndPoint, issueID)
	issue.Fields.Description = gitURL
	marshalledIssue, err := json.Marshal(issue)
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
