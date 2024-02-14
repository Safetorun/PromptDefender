Feature: Keep tests

  Scenario: Sending a request for a suspicious user should true for suspicious response
    Given I create a suspicious user with the user_id this-is-a-suspicious-user
    And I send a request to moat
    And the user_id is this-is-a-suspicious-user
    And the request is hello world
    And request is sent to moat
    And Suspicious user response should be true
