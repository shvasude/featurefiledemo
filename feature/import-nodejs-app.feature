@e2e
Feature: Importing nodejs application

    @golden-path
    Scenario: Install master version of the Service Binding Operator via OperatorSource
        Given cluster is available
        And active project is set to 'service-binding-demo'
        
        When "oc new-app nodejs~"<nodeJSApp>"--name "<appName> command is executed
        Then buildconfig is created with the name <appName>
        And build is created with the name <appName>        
        And build gets to a "Complete" status within "5" minutes

        And pod with name starting with <appName> is created 
        And pod with name starting with <appName> gets to the "Succeeded" state within "2" minutes
        And get the name of the deployment config matched to build config

        When deployment is used instead of deployment config
        And when "oc expose svc/"<buildconfig>"--name="<buildconfig> command is executed
        Then route is created when "oc get route"<buildconfig> command is executed
        And "N/A" is seen as a header in the application URL given by the route
        
        Examples:
        | nodeJSApp                                         | appName               | 
        | https://github.com/pmacik/nodejs-rest-http-crud   | nodejs-rest-http-crud | 