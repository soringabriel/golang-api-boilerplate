package user_requests

import (
	"encoding/json"
	"errors"
	"net/http"
)

type UpdateFilterRequestParams struct {
	Email *string `json:"email"`
}

type UpdateSetRequestParams struct {
	Name *string `json:"name"`
}

type UpdateRequestParams struct {
	Filter *UpdateFilterRequestParams `json:"filter"`
	Set    *UpdateSetRequestParams    `json:"set"`
}

type UpdateRequest struct {
	Method string              `json:"method"`
	Params UpdateRequestParams `json:"params"`
}

func (request *UpdateRequest) Defaults() {
	// No defaults to set
}

func (request *UpdateRequest) Read(r *http.Request) {
	// Set request method
	request.Method = r.Method

	// Unmarshal JSON body
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&request.Params)

	request.Defaults()
}

func (request *UpdateRequest) Validate() error {
	if request.Method != http.MethodPut {
		return errors.New("method must be PUT")
	}
	if request.Params.Filter == nil || request.Params.Filter.Email == nil {
		return errors.New("email is required")
	}
	return nil
}

func (request *UpdateRequest) GetMethod() string {
	return request.Method
}

func (request *UpdateRequest) GetParams() []byte {
	params, _ := json.Marshal(request.Params)
	return params
}
