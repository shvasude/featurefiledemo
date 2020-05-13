Feature: Sample Demo
  I need to be able to use feature file to upload this to Polarion  

  Scenario: Searching for "first set of operation"
    Given I go to "https://google.com"
    When I search for "gmail"
    Then I should see "Gmail"

  Scenario: Select GMAIL link
    Given I see Gmail from the above scenario
    When I select the GMAIL link from the list of searches
    Then I should navigate to GMAIL website
