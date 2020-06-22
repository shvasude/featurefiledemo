@e2e
Feature: Install Service Binding Operator

    @golden-path
    Scenario: Install master version of the Service Binding Operator via OperatorSource
        Given cluster is available
        When <CMD_SRC> command is executed 
        Then install plan owned by <operatorName> subscription is created in the <operatorsNS> namespace
        And install plan owned by <operatorName> subscription gets to the 'Complete' state within 3 minutes
        And pod with name starting with <operatorName> is created in the <operatorsNS> namespace
        And pod with name starting with <operatorName> gets to the 'Running' state within 3 minutes

        Examples:
        | CMD_SRC                                       | operatorsNS           | operatorName              | 
        | make install-service-binding-operator-master  | openshift-operators   | service-binding-operator  |