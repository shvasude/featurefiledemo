@e2e
Feature: Uninstall Backing Service Operator

    @golden-path
    Scenario: Uninstalling backing service operator
        Given cluster is available
        When <CMD> command is executed to "uninstall backing service operator"
        Then backing service operator subscription is deleted with the message "deleted" for <pkgManifest>
        And backing service operator source is deleted with the message "deleted" for <pkgManifest>
        And <pkgManifest> not found for <operatorsNS> namespace with the message "not found"
        And operator subscription for <operatorsNS> namespace "not found"
        And pod that was running <pkgManifest> backing service operator "not found" that takes "5" minutes to update

        Examples:
        | CMD                               | pkgManifest   | operatorsNS |
        | make uninstall-backing-db-operator| db-operators  | openshift-operators  |