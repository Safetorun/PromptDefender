Feature: Keep tests

  Scenario: Sending a request for a suspicious user should true for suspicious response
    Given I create a suspicious user with the user_id this-is-a-suspicious-user
    And I send a request to wall
    And the user_id is this-is-a-suspicious-user
    And the request is hello world
    And request is sent to wall
    And Suspicious user response should be true

Scenario: Sending a request for a suspicious user should be false for suspicious response if doesnt exist
  Given I send a request to wall
  And the user_id is this-is-a-blah-blah-user
  And the request is hello world
  And request is sent to wall
  And Suspicious user response should be false
