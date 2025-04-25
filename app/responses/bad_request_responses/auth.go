package bad_request_responses

import (
	"net/http"

	"api/app/responses"
)

var MissingAuthResponse = responses.Response{
	StatusCode: http.StatusUnauthorized,
	Body: responses.ResponseBody{
		Status: responses.RESPONSE_STATUS_ERROR,
		Error:  "Unauthorized",
	},
	Headers: map[string]string{
		"Content-Type": "application/json",
	},
}

var WrongAuthResponse = responses.Response{
	StatusCode: http.StatusForbidden,
	Body: responses.ResponseBody{
		Status: responses.RESPONSE_STATUS_ERROR,
		Error:  "Forbidden",
	},
	Headers: map[string]string{
		"Content-Type": "application/json",
	},
}
