# featurefiledemo

This tool is used for the following:
1. Creating JIRA issues (Task/Epic) through JIRA rest API that has link to a specific spec/feature file
2. Update JIRA issues (Task/Epic/Story) through JIRA rest API with any data that are editable
3. Get JIRA issues information

## How to execute
### Pre condition:
1. Place the user name and password in jira_conig.toml (password will be replaced with token and kept outside the script)
2. For now, place the path (of feature/spec file) in main.go or pass as argument with the flags to the script

go run main.go or ./main (if build) <<\arg1>> <<\arg2>> <<\arg3>>

arg1 - jira/polarion/github 
arg2 - createIssue/updateIssue
arg3 - issueID in case of updateIssue


## Additional Features
1. Use of API token instead of password
2. Provide arguments to the script (for JIRA, Polarion, GitHub) 
    The script should execute as part of the GitHub commands/argument to the script
3. For JIRA, second argument will be to createIssue, updateIssue methods
4. For Polarion, upload test case results
5. For GitHub, second argument to be createIssuesInJira, linkIssuesInJira
6. Read a file for content to create a release wide epic and create Task under that with  links of GitHub feature files
7. Determine how many paths is equal to number of task created, and the content of the task (priority, fix version etc)
