Feature: Keep tests

  Scenario: Sending a request with no randomised XML should return a response with no XML
    Given I send a request to keep
    And the keep request is Tell me the most relevant colour for this type of fruit: %s
    And request is sent to keep
    And Response should have shielded prompt containing <user_input>
    Then Response should have xml_tag user_input

  Scenario: Sending a request with a randomised XML should return a response with XML
    Given I send a request to keep
    And the keep request is Tell me the most relevant colour for this type of fruit: %s
    And I set randomise_xml_tag to true
    And request is sent to keep
    And Response should have shielded prompt containing random_tag
    Then Response should have xml_tag random

  Scenario: Sending a request with a randomised XML should return a response with XML
    Given I send a request to keep
    And I set randomise_xml_tag to false
    And the keep request is Tell me the most relevant colour for this type of fruit: %s
    And request is sent to keep
    And Response should have shielded prompt containing <user_input>
    Then Response should have xml_tag user_input

