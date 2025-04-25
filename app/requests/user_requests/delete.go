package user_requests

import (
	"encoding/json"
	"errors"
	"net/http"
)

type DeleteRequestParams struct {
	Email *string `json:"email"`
}

type DeleteRequest struct {
	Method string              `json:"method"`
	Params DeleteRequestParams `json:"params"`
}

func (request *DeleteRequest) Defaults() {
	// No defaults to set
}

func (request *DeleteRequest) Read(r *http.Request) {
	// Set request method
	request.Method = r.Method

	// Unmarshal JSON body
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&request.Params)

	request.Defaults()
}

func (request *DeleteRequest) Validate() error {
	if request.Method != http.MethodDelete {
		return errors.New("method must be DELETE")
	}
	if request.Params.Email == nil {
		return errors.New("email is required")
	}
	return nil
}

func (request *DeleteRequest) GetMethod() string {
	return request.Method
}

func (request *DeleteRequest) GetParams() []byte {
	params, _ := json.Marshal(request.Params)
	return params
}
