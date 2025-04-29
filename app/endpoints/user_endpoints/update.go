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

var UpdateUserEndpoint = &endpoints.Endpoint{
	Path: "/user/update",
	Middlewares: []middlewares.Middleware{
		middlewares.AuthMiddlewareFactory(),
	},
	HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
		// Read and validate request
		request := user_requests.UpdateRequest{}
		request.Read(r)
		if err := request.Validate(); err != nil {
			bad_request_responses.BadRequestResponse(&request, err).Write(w)
			return
		}

		// Update data
		filter := bson.D{{Key: "email", Value: request.Params.Filter.Email}}
		var update bson.D
		if request.Params.Set != nil && request.Params.Set.Name != nil {
			update = bson.D{
				{Key: "$set", Value: bson.D{
					{Key: "name", Value: request.Params.Set.Name},
				}},
			}
		} else {
			update = bson.D{
				{Key: "$unset", Value: bson.D{
					{Key: "name", Value: ""},
				}},
			}
		}
		updateResult, updateResultErr := databases.MongodbDatabase.UpdateOne(user_model.CollectionName, filter, update)

		// Check for errors
		if updateResultErr != nil {
			bad_request_responses.StatusInternalServerError(&request, updateResultErr).Write(w)
			return
		}

		// Return response
		response := responses.Response{
			Body: responses.ResponseBody{
				Request: &request,
				Status:  responses.RESPONSE_STATUS_SUCCESS,
				Response: map[string]interface{}{
					"updated": updateResult.MatchedCount > 0,
				},
			},
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		response.Write(w)
	},
}
