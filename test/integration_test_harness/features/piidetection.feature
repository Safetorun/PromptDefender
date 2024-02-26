Feature: Pii Detection

  Scenario: PII detection is on and PII is sent
    Given I send a request to wall
    When I set PII detection to true
    And the request is Franklin Roosevelt, 13 George Street, GL55 4PE
    And request is sent to wall
    Then Response should have PII detected set to true

  Scenario: PII detection is on, but no PII is sent
    Given I send a request to wall
    When I set PII detection to false
    And the request is benign and not containing PII
    And request is sent to wall
    Then Response should have PII detected set to false

  Scenario: PII detection is off and PII is sent
    Given I send a request to wall
    When I set PII detection to true
    And the request is Franklin Roosevelt, 13 George Street, GL55 4PE
    And request is sent to wall
    Then Response should have PII detected set to true

  Scenario: PII detection is off and PII is not sent
    Given I send a request to wall
    When I set PII detection to false
    And the request is benign and not containing PII
    And request is sent to wall
    Then Response should have PII detected set to nil