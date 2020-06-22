@e2e
Feature: create a service binding request 

    @golden-path
    Scenario: Create a service binding request to bind backing db service and nodejs application
        
        Given cluster is available

        When <CMD> command is executed
        Then "service-binding request" is created
        And 'service-binding-request' gets to the 'True' state
        And annotation is matched with <appNS>, <sbr>, <appName>, <expKind>, <dbName>, <expResource>
        Then route is created when "oc get route"<buildconfig> command is executed
        And <dbName> is seen as a header in the application URL given by the route
        And secret will have an name of the service binding request <sbr>
        And deployment env should have <env>
        And deployment env From should be <sbr> 
        
        Examples:
        | CMD                                   | appNS                 | sbr                       | appName               | expKind   | dbName    | expResource   | env                   |
        | make create-service-binding-request   | service-binding-demo  | service-binding-request   | nodejs-rest-http-crud | Database  | db-demo   | deployments   | ServiceBindingOperator|