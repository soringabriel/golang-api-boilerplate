package user_endpoints

import (
	"net/http"

	"api/app/endpoints"
	"api/app/middlewares"
	"api/app/models/user_model"
	"api/app/requests/user_requests"
	"api/app/responses"
	"api/app/responses/bad_request_responses"
	"api/databases"
)

var CreateUserEndpoint = &endpoints.Endpoint{
	Path:        "/user/create",
	Middlewares: []middlewares.Middleware{middlewares.AuthMiddleware},
	HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
		// Read and validate request
		request := user_requests.CreateRequest{}
		request.Read(r)
		if err := request.Validate(); err != nil {
			bad_request_responses.BadRequestResponse(&request, err).Write(w)
			return
		}

		// Insert data
		newModel := user_model.User{
			Email: *request.Params.Email,
			Name:  request.Params.Name,
		}
		_, insertResultErr := databases.MongodbDatabase.InsertOne(user_model.CollectionName, newModel.ToBson())

		// Check for errors
		if insertResultErr != nil {
			bad_request_responses.StatusInternalServerError(&request, insertResultErr).Write(w)
			return
		}

		// Return response
		response := responses.Response{
			Body: responses.ResponseBody{
				Request: &request,
				Status:  responses.RESPONSE_STATUS_SUCCESS,
				Response: map[string]interface{}{
					"created": true,
				},
			},
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		response.Write(w)
	},
}
