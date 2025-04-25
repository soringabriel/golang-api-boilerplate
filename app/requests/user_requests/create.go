package user_requests

import (
	"encoding/json"
	"errors"
	"net/http"
)

type CreateRequestParams struct {
	Email *string `json:"email"`
	Name  *string `json:"name"`
}

type CreateRequest struct {
	Method string              `json:"method"`
	Params CreateRequestParams `json:"params"`
}

func (request *CreateRequest) Defaults() {
	// No defaults to set
}

func (request *CreateRequest) Read(r *http.Request) {
	// Set request method
	request.Method = r.Method

	// Unmarshal JSON body
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&request.Params)

	request.Defaults()
}

func (request *CreateRequest) Validate() error {
	if request.Method != http.MethodPost {
		return errors.New("method must be POST")
	}
	if request.Params.Email == nil {
		return errors.New("email is required")
	}
	return nil
}

func (request *CreateRequest) GetMethod() string {
	return request.Method
}

func (request *CreateRequest) GetParams() []byte {
	params, _ := json.Marshal(request.Params)
	return params
}
