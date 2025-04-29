package bad_request_responses

import (
	"net/http"

	"api/app/responses"
)

var RateLimitResponse = responses.Response{
	StatusCode: http.StatusTooManyRequests,
	Body: responses.ResponseBody{
		Status: responses.RESPONSE_STATUS_ERROR,
		Error:  "Too many requests",
	},
	Headers: map[string]string{
		"Content-Type": "application/json",
	},
}
