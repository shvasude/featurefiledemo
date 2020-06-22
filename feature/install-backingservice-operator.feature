@e2e
Feature: Install Backing Service Operator

    @golden-path
    Scenario: Install Backing service operator via OperatorSource
        
        Given cluster is available        
        When <CMD-SRC> command is executed to install backing service operator source        
        Then there is an entry for <pkgManifest> in package manifest
        When <CMD-SUB> command is executed to install backing service operator subscription
        Then install plan owned by <pkgManifest> subscription is created in the <operatorsNS> namespace        
        And install plan owned by <pkgManifest> subscription gets to the 'Complete' state within '3' minutes
        And pod with name starting with <pkgManifest> is created in the <operatorsNS> namespace
        And pod with name starting with <pkgManifest> gets to the 'Running' state within '3' minutes
        And there is an entry in the crd for 'databases.postgresql.baiju.dev'

        Examples:
        | CMD-SRC                                           | CMD-SUB                                       | pkgManifest  | operatorsNS        |
        | make install-backing-db-operator-source           | make install-backing-db-operator-subscription | db-operators | openshift-operators|