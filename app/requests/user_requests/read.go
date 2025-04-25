package user_requests

import (
	"api/helpers"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/schema"
)

type ReadRequestParams struct {
	Limit *int64 `json:"limit" schema:"limit"`
	Skip  *int64 `json:"skip" schema:"skip"`
}

type ReadRequest struct {
	Method string            `json:"method"`
	Params ReadRequestParams `json:"params"`
}

func (request *ReadRequest) Defaults() {
	if request.Params.Limit == nil {
		limit := int64(10)
		request.Params.Limit = &limit
	}
	if request.Params.Skip == nil {
		skip := int64(0)
		request.Params.Skip = &skip
	}
}

func (request *ReadRequest) Read(r *http.Request) {
	// Set request method
	request.Method = r.Method

	// Get filter_params from query parameters
	decoder := schema.NewDecoder()
	decoder.Decode(&request.Params, r.URL.Query())

	request.Defaults()
}

func (request *ReadRequest) Validate() error {
	if request.Method != http.MethodGet {
		return errors.New("method must be GET")
	}
	if *request.Params.Limit < 0 || *request.Params.Limit > 100 {
		return errors.New("limit must be between 0 and 100")
	}
	if *request.Params.Skip < 0 {
		return errors.New("skip must be greater than or equal to 0")
	}
	return nil
}

func (request *ReadRequest) GetMethod() string {
	return request.Method
}

func (request *ReadRequest) GetParams() []byte {
	params, _ := json.Marshal(helpers.StructToMapString(request.Params))
	return params
}
