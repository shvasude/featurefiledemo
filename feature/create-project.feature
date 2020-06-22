@e2e
Feature: Creating a project in OCP through cli

    @golden-path
    Scenario: Install master version of the Service Binding Operator via OperatorSource
        Given cluster is available
        
        When <CMD> command is executed
        Then message with <appNS> is received
        And <appNS> name is added as part of the metadata of <appNS> namespace
        
        And Project <appNS> is created with the "Active"
        And active project is set to <appNS>

        Examples:
        | CMD                   |appNS                 | 
        | make create-project   |service-binding-demo  | 