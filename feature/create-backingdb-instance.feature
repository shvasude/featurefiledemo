@e2e
Feature: Create a backing db instance

    @golden-path
    Scenario: Create a backing db instance
        Given cluster is available        
        When <CMD> command is executed        
        Then db instance with the name <dbName> is created
        And db instance gets a valid connection IP 
        And pod with name starting with <dbName> is created

        Examples:
        | CMD                               | dbName | 
        | make create-backing-db-instance   | db-demo|