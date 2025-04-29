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

	"go.mongodb.org/mongo-driver/bson"
)

var DeleteUserEndpoint = &endpoints.Endpoint{
	Path: "/user/delete",
	Middlewares: []middlewares.Middleware{
		middlewares.AuthMiddlewareFactory(),
	},
	HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
		// Read and validate request
		request := user_requests.DeleteRequest{}
		request.Read(r)
		if err := request.Validate(); err != nil {
			bad_request_responses.BadRequestResponse(&request, err).Write(w)
			return
		}

		// Delete data
		filter := bson.D{{Key: "email", Value: request.Params.Email}}
		deleteResult, deleteResultErr := databases.MongodbDatabase.DeleteOne(user_model.CollectionName, filter)

		// Check for errors
		if deleteResultErr != nil {
			bad_request_responses.StatusInternalServerError(&request, deleteResultErr).Write(w)
			return
		}

		// Return response
		response := responses.Response{
			Body: responses.ResponseBody{
				Request: &request,
				Status:  responses.RESPONSE_STATUS_SUCCESS,
				Response: map[string]interface{}{
					"deleted": deleteResult.DeletedCount > 0,
				},
			},
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		response.Write(w)
	},
}
