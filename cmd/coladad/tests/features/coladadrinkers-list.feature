@drinkers @list
Feature: API Clients are able to list information about people who drink coladas

    # Success Scenarios
    Scenario: API clients call /v1/drinkers
      When a request is made to "GET" "/v1/drinkers"
      Then the response should be a 200 status code
      And the response body should contain a result with at least 1 item
      