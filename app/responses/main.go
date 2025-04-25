package responses

import (
	"api/app/requests"
	"encoding/json"
	"net/http"
)

const RESPONSE_STATUS_SUCCESS = "success"
const RESPONSE_STATUS_ERROR = "error"

type ResponseBody struct {
	Request  requests.Request `json:"request"`
	Status   string           `json:"status"`
	Error    string           `json:"error"`
	Response interface{}      `json:"response"`
}

type Response struct {
	StatusCode int               `json:"status_code"`
	Headers    map[string]string `json:"headers"`
	Body       ResponseBody      `json:"body"`
}

func (response *Response) Write(w http.ResponseWriter) {
	for key, value := range response.Headers {
		w.Header().Set(key, value)
	}
	if response.StatusCode != 0 {
		w.WriteHeader(response.StatusCode)
	}
	json.NewEncoder(w).Encode(response.Body)
}
