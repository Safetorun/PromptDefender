Feature: Banned users

  Scenario: I should be able to create, list and delete a user
    Given I create a suspicious user with the user_id user1234
    And I retrieve a list of all users
    Then The list should contain the user_id user1234
    And I delete the user with the user_id user1234
    And I retrieve a list of all users
    Then The list should not contain the user_id user1234

  Scenario: I should get 404 if a user doesnt exist
    Given I retrieve the user with the user_id user1234
    Then I should get a 404 error