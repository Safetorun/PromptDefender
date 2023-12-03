Feature: Keep tests

  Scenario: Sending a request with no randomised XML should return a response with no XML
    Given I send a request to keep
    And request is sent to moat
    Then Response should have xml_tag user_input

  Scenario: Sending a request with a randomised XML should return a response with XML
    Given I send a request to keep
    And I set randomise_xml_tag to true
    And request is sent to moat
    Then Response should have a random xml_tag

  Scenario: Sending a request with a randomised XML should return a response with XML
    Given I send a request to keep
    And I set randomise_xml_tag to false
    And request is sent to moat
    Then Response should have xml_tag user_input

