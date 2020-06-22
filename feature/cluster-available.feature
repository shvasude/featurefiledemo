@e2e
Feature: Check cluster availability

    @golden-path
    Scenario: Cluster is available
        Given KUBECONFIG variable is set properly
        And 'oc status' command is executed
        Then exit code of the command is '0'
        And cluster is available

    @negative-path
    Scenario: Cluster is not available
        Given KUBECONFIG variable is set to a wrong cluster
        And 'oc status' command is executed
        Then exit code of the command is not '0'
        And cluster is not available