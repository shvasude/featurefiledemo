Feature: Sample Demo
  I need to be able to use feature file to upload this to Polarion  

  Scenario: Searching for "first set of operation"
    Given I go to "https://google.com"
    When I search for "gmail"
    Then I should see "Gmail"
