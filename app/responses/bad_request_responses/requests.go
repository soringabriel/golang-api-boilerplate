package bad_request_responses

import (
	"api/app/requests"
	"api/app/responses"
	"net/http"
)

func BadRequestResponse(request requests.Request, err error) *responses.Response {
	return &responses.Response{
		StatusCode: http.StatusBadRequest,
		Body: responses.ResponseBody{
			Request: request,
			Status:  responses.RESPONSE_STATUS_ERROR,
			Error:   err.Error(),
		},
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}
}

func StatusInternalServerError(request requests.Request, err error) *responses.Response {
	return &responses.Response{
		StatusCode: http.StatusInternalServerError,
		Body: responses.ResponseBody{
			Request: request,
			Status:  responses.RESPONSE_STATUS_ERROR,
			Error:   err.Error(),
		},
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}
}
