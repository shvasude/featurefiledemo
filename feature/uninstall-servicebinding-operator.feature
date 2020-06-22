@e2e
Feature: Uninstall Service Binding Operator

    @golden-path
    Scenario: Uninstalling service binding operator
        Given cluster is available
        When <CMD> command is executed to "uninstall service binding operator"
        Then service binding operator subscription is deleted with the message "deleted" for <pkgManifest> for <operatorName>
        And backing service operator source is deleted with the message "deleted" for <pkgManifest>
        And <pkgManifest> not found for <operatorsNS> namespace with the message "not found"
        And operator subscription for <operatorsNS> namespace "not found"
        And pod that was running <pkgManifest> backing service operator "not found" that takes "5" minutes to update

        Examples:
        | CMD                               | pkgManifest   | operatorsNS | operatorName |
        | make uninstall-service-binding-operator-master| db-operators  | openshift-operators  | service-binding-operator |