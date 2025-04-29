package tests

import (
	"fmt"
	"testing"

	"api/app/endpoints/user_endpoints"
	"api/app/requests/user_requests"
	"api/app/responses/bad_request_responses"
	"api/helpers"
)

func TestMissingAuth(t *testing.T) {
	t.Parallel()

	err := helpers.TestRequest(
		fmt.Sprintf(
			"http://%s%s",
			helpers.GetEnvVariable("API_IP_PORT"),
			user_endpoints.ReadUserEndpoint.Path,
		),
		&user_requests.ReadRequest{},
		map[string]string{},
		bad_request_responses.MissingAuthResponse,
	)
	if err != nil {
		t.Fatal(err)
	}
}

func TestWrongAuth(t *testing.T) {
	t.Parallel()

	err := helpers.TestRequest(
		fmt.Sprintf(
			"http://%s%s",
			helpers.GetEnvVariable("API_IP_PORT"),
			user_endpoints.ReadUserEndpoint.Path,
		),
		&user_requests.ReadRequest{},
		map[string]string{
			"Accept":        "application/json",
			"Authorization": "Bearer <WRONG_AUTH_TOKEN>",
		},
		bad_request_responses.WrongAuthResponse,
	)
	if err != nil {
		t.Fatal(err)
	}
}
