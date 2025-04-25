package user_endpoints

import (
	"net/http"
	"sync"

	"api/app/endpoints"
	"api/app/middlewares"
	"api/app/models/user_model"
	"api/app/requests/user_requests"
	"api/app/responses"
	"api/app/responses/bad_request_responses"
	"api/databases"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ReadUserEndpoint = &endpoints.Endpoint{
	Path:        "/user/read",
	Middlewares: []middlewares.Middleware{middlewares.AuthMiddleware},
	HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
		// Read and validate request
		request := user_requests.ReadRequest{}
		request.Read(r)
		if err := request.Validate(); err != nil {
			bad_request_responses.BadRequestResponse(&request, err).Write(w)
			return
		}

		// Read data (concurrent queries for results and total results)
		results := []*user_model.User{}
		var totalResults int64
		var readResultErr error
		var totalResultsErr error
		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			defer wg.Done()
			_, readResultErr = databases.MongodbDatabase.Read(
				user_model.CollectionName,
				bson.M{},
				&options.FindOptions{
					Limit: request.Params.Limit,
					Skip:  request.Params.Skip,
				},
				&results,
			)
		}()
		go func() {
			defer wg.Done()
			totalResults, totalResultsErr = databases.MongodbDatabase.Count(user_model.CollectionName, bson.M{})
		}()
		wg.Wait()

		// Check for errors
		if readResultErr != nil {
			bad_request_responses.StatusInternalServerError(&request, readResultErr).Write(w)
			return
		}
		if totalResultsErr != nil {
			bad_request_responses.StatusInternalServerError(&request, totalResultsErr).Write(w)
			return
		}

		// Return response
		response := responses.Response{
			Body: responses.ResponseBody{
				Request: &request,
				Status:  responses.RESPONSE_STATUS_SUCCESS,
				Response: map[string]interface{}{
					"results":       results,
					"total_results": totalResults,
				},
			},
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		response.Write(w)
	},
}
