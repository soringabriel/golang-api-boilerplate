package tests

import (
	"api/app/endpoints/user_endpoints"
	"api/app/models/user_model"
	"api/app/requests/user_requests"
	"api/app/responses"
	"api/app/responses/bad_request_responses"
	"api/helpers"
	"errors"
	"fmt"
	"net/http"
	"testing"
)

func TestUserEndpointsValidation(t *testing.T) {
	t.Parallel()

	// Create
	createRequest := user_requests.CreateRequest{Method: http.MethodPost}
	createRequest.Defaults()
	err := helpers.TestRequest(
		fmt.Sprintf(
			"http://%s%s",
			helpers.EnvVariable("API_IP_PORT"),
			user_endpoints.CreateUserEndpoint.Path,
		),
		&createRequest,
		map[string]string{
			"Accept":        "application/json",
			"Authorization": fmt.Sprintf("Bearer %s", helpers.EnvVariable("AUTH_TOKEN")),
		},
		*bad_request_responses.BadRequestResponse(&createRequest, errors.New("email is required")),
	)
	if err != nil {
		t.Fatal(err)
	}

	// Read
	limit := int64(1000)
	readRequest := user_requests.ReadRequest{Method: http.MethodGet, Params: user_requests.ReadRequestParams{Limit: &limit}}
	readRequest.Defaults()
	err = helpers.TestRequest(
		fmt.Sprintf(
			"http://%s%s",
			helpers.EnvVariable("API_IP_PORT"),
			user_endpoints.ReadUserEndpoint.Path,
		),
		&readRequest,
		map[string]string{
			"Accept":        "application/json",
			"Authorization": fmt.Sprintf("Bearer %s", helpers.EnvVariable("AUTH_TOKEN")),
		},
		*bad_request_responses.BadRequestResponse(&readRequest, errors.New("limit must be between 0 and 100")),
	)
	if err != nil {
		t.Fatal(err)
	}

	// Update
	updateRequest := user_requests.UpdateRequest{Method: http.MethodPut}
	updateRequest.Defaults()
	err = helpers.TestRequest(
		fmt.Sprintf(
			"http://%s%s",
			helpers.EnvVariable("API_IP_PORT"),
			user_endpoints.UpdateUserEndpoint.Path,
		),
		&updateRequest,
		map[string]string{
			"Accept":        "application/json",
			"Authorization": fmt.Sprintf("Bearer %s", helpers.EnvVariable("AUTH_TOKEN")),
		},
		*bad_request_responses.BadRequestResponse(&updateRequest, errors.New("email is required")),
	)
	if err != nil {
		t.Fatal(err)
	}

	// Delete project
	deleteRequest := user_requests.DeleteRequest{Method: http.MethodDelete}
	deleteRequest.Defaults()
	err = helpers.TestRequest(
		fmt.Sprintf(
			"http://%s%s",
			helpers.EnvVariable("API_IP_PORT"),
			user_endpoints.DeleteUserEndpoint.Path,
		),
		&deleteRequest,
		map[string]string{
			"Accept":        "application/json",
			"Authorization": fmt.Sprintf("Bearer %s", helpers.EnvVariable("AUTH_TOKEN")),
		},
		*bad_request_responses.BadRequestResponse(&deleteRequest, errors.New("email is required")),
	)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUserEndpoints(t *testing.T) {
	t.Parallel()

	// User
	user := user_model.User{Email: "test@test.com"}

	// Create user
	createRequest := user_requests.CreateRequest{
		Method: http.MethodPost,
		Params: user_requests.CreateRequestParams{Email: &user.Email},
	}
	createRequest.Defaults()
	err := helpers.TestRequest(
		fmt.Sprintf(
			"http://%s%s",
			helpers.EnvVariable("API_IP_PORT"),
			user_endpoints.CreateUserEndpoint.Path,
		),
		&createRequest,
		map[string]string{
			"Accept":        "application/json",
			"Authorization": fmt.Sprintf("Bearer %s", helpers.EnvVariable("AUTH_TOKEN")),
		},
		responses.Response{
			StatusCode: http.StatusOK,
			Body: responses.ResponseBody{
				Request: &createRequest,
				Status:  responses.RESPONSE_STATUS_SUCCESS,
				Response: map[string]interface{}{
					"created": true,
				},
			},
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	// Read user
	readRequest := user_requests.ReadRequest{Method: http.MethodGet}
	readRequest.Defaults()
	err = helpers.TestRequest(
		fmt.Sprintf(
			"http://%s%s",
			helpers.EnvVariable("API_IP_PORT"),
			user_endpoints.ReadUserEndpoint.Path,
		),
		&readRequest,
		map[string]string{
			"Accept":        "application/json",
			"Authorization": fmt.Sprintf("Bearer %s", helpers.EnvVariable("AUTH_TOKEN")),
		},
		responses.Response{
			StatusCode: http.StatusOK,
			Body: responses.ResponseBody{
				Request: &readRequest,
				Status:  responses.RESPONSE_STATUS_SUCCESS,
				Response: map[string]interface{}{
					"results":       []user_model.User{user},
					"total_results": int64(1),
				},
			},
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	// Update user
	newName := "test"
	filterParams := user_requests.UpdateFilterRequestParams{Email: &user.Email}
	setParams := user_requests.UpdateSetRequestParams{Name: &newName}
	updateRequest := user_requests.UpdateRequest{
		Method: http.MethodPut,
		Params: user_requests.UpdateRequestParams{
			Filter: &filterParams,
			Set:    &setParams,
		},
	}
	updateRequest.Defaults()
	err = helpers.TestRequest(
		fmt.Sprintf(
			"http://%s%s",
			helpers.EnvVariable("API_IP_PORT"),
			user_endpoints.UpdateUserEndpoint.Path,
		),
		&updateRequest,
		map[string]string{
			"Accept":        "application/json",
			"Authorization": fmt.Sprintf("Bearer %s", helpers.EnvVariable("AUTH_TOKEN")),
		},
		responses.Response{
			StatusCode: http.StatusOK,
			Body: responses.ResponseBody{
				Request: &updateRequest,
				Status:  responses.RESPONSE_STATUS_SUCCESS,
				Response: map[string]interface{}{
					"updated": true,
				},
			},
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	user.Name = &newName

	// Read user
	err = helpers.TestRequest(
		fmt.Sprintf(
			"http://%s%s",
			helpers.EnvVariable("API_IP_PORT"),
			user_endpoints.ReadUserEndpoint.Path,
		),
		&readRequest,
		map[string]string{
			"Accept":        "application/json",
			"Authorization": fmt.Sprintf("Bearer %s", helpers.EnvVariable("AUTH_TOKEN")),
		},
		responses.Response{
			StatusCode: http.StatusOK,
			Body: responses.ResponseBody{
				Request: &readRequest,
				Status:  responses.RESPONSE_STATUS_SUCCESS,
				Response: map[string]interface{}{
					"results":       []user_model.User{user},
					"total_results": int64(1),
				},
			},
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	// Delete user
	deleteRequest := user_requests.DeleteRequest{
		Method: http.MethodDelete,
		Params: user_requests.DeleteRequestParams{Email: &user.Email},
	}
	deleteRequest.Defaults()
	err = helpers.TestRequest(
		fmt.Sprintf(
			"http://%s%s",
			helpers.EnvVariable("API_IP_PORT"),
			user_endpoints.DeleteUserEndpoint.Path,
		),
		&deleteRequest,
		map[string]string{
			"Accept":        "application/json",
			"Authorization": fmt.Sprintf("Bearer %s", helpers.EnvVariable("AUTH_TOKEN")),
		},
		responses.Response{
			StatusCode: http.StatusOK,
			Body: responses.ResponseBody{
				Request: &deleteRequest,
				Status:  responses.RESPONSE_STATUS_SUCCESS,
				Response: map[string]interface{}{
					"deleted": true,
				},
			},
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	// Delete user
	deleteRequest.Defaults()
	err = helpers.TestRequest(
		fmt.Sprintf(
			"http://%s%s",
			helpers.EnvVariable("API_IP_PORT"),
			user_endpoints.DeleteUserEndpoint.Path,
		),
		&deleteRequest,
		map[string]string{
			"Accept":        "application/json",
			"Authorization": fmt.Sprintf("Bearer %s", helpers.EnvVariable("AUTH_TOKEN")),
		},
		responses.Response{
			StatusCode: http.StatusOK,
			Body: responses.ResponseBody{
				Request: &deleteRequest,
				Status:  responses.RESPONSE_STATUS_SUCCESS,
				Response: map[string]interface{}{
					"deleted": false,
				},
			},
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	// Read user
	err = helpers.TestRequest(
		fmt.Sprintf(
			"http://%s%s",
			helpers.EnvVariable("API_IP_PORT"),
			user_endpoints.ReadUserEndpoint.Path,
		),
		&readRequest,
		map[string]string{
			"Accept":        "application/json",
			"Authorization": fmt.Sprintf("Bearer %s", helpers.EnvVariable("AUTH_TOKEN")),
		},
		responses.Response{
			StatusCode: http.StatusOK,
			Body: responses.ResponseBody{
				Request: &readRequest,
				Status:  responses.RESPONSE_STATUS_SUCCESS,
				Response: map[string]interface{}{
					"results":       []user_model.User{},
					"total_results": int64(0),
				},
			},
		},
	)
	if err != nil {
		t.Fatal(err)
	}
}
