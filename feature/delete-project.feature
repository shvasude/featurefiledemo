@e2e
Feature: Delete Project

    @golden-path
    Scenario: Deleting a project in cluster
        Given cluster is available
        When <CMD> command is executed to "delete the project"
        Then <appNS> project is deleted with "deleted" message
        And once <appNS> project is deleted, gives the message <appNS> "not found" within "5" minutes
        
        Examples:
        | CMD                   | appNS | 
        | make delete-project   | service-binding-demo  |