@e2e
Feature: Service Binding
    As a user of service binding Service Binding Operator
    I want to bind applications with backing services

    Background: 
    Given Service Binding Operator is installed

    @example @e2e
    Scenario: To bind an imported nodejs app to PostgreSQL database
        
        Given PostgreSQL db operator is installed
        And namespace is created 
        And Nodejs application is imported
        And Postgres database instance is created
        When Nodejs application is connected to db using Service Binding Request
        Then Nodejs application is up and running connected to the Postgres database