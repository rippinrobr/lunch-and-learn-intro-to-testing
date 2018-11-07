@drawings
Feature: API Clients want a list of colada drinkers and the ability to select a barista and cleaner

  Scenario: API clients want to see the previously selected barista, cleaner pairing
    When a request is made to "GET" "/v1/drawings/previous"
    Then the response should be a 200 status code
    And the response body should return a single item
    And the response body should contain the field "barista"
    And the response body should contain the field "cleaner"

  
  Scenario: API client wants to generate a new paring of baristas and cleaners
    When a request is made to "POST" "/v1/drawings"
    Then the response should be a 200 status code
    And the response body should return a single item
    And the response body should contain the field "barista"
    And the response body should contain the field "baristaImg"
    And the response body should contain the field "baristaId"
    And the response body should contain the field "cleaner"
    And the response body should contain the field "cleanerImg"
    And the response body should contain the field "cleanerId"

  Scenario: Ensure that the generated paring of barista and cleaner are not the  same as the ones in the previous drawing
    When a request is made to "GET" "/v1/drawings/previous"
    Then the response should be a 200 status code
    And the response body should return a single item
    When a request is made to "POST" "/v1/drawings"
    Then the response should be a 200 status code
    And the "barista" field does not equal the previous "barista"
    And the "cleaner" field does not equal the previous "cleaner"
    And the "cleaner" field does not equal the previous "barista"