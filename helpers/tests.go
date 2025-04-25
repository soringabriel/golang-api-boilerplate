package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"api/app/requests"
	"api/app/responses"
)

func TestRequest(
	url string,
	request requests.Request,
	requestHeaders map[string]string,
	expectedResponse responses.Response,
) error {
	// Make request
	method := request.GetMethod()
	httpRequestParams := HttpRequestParams{
		Method:  method,
		Url:     url,
		Headers: requestHeaders,
	}
	if method == http.MethodGet {
		fmt.Println(string(request.GetParams()))
		err := json.Unmarshal(request.GetParams(), &httpRequestParams.QueryParams)
		if err != nil {
			return fmt.Errorf("error unmarshalling params: %v", err)
		}
	} else {
		httpRequestParams.RequestBody = request.GetParams()
	}
	resp, respBody, err := HttpRequest(httpRequestParams)
	if err != nil {
		return fmt.Errorf("error making request: %v", err)
	}

	// Check response
	if resp.StatusCode != expectedResponse.StatusCode {
		return fmt.Errorf("unexpected status code: %d. Expected: %d", resp.StatusCode, expectedResponse.StatusCode)
	}
	expectedBody, _ := json.Marshal(expectedResponse.Body)
	if strings.TrimSpace(string(respBody)) != strings.TrimSpace(string(expectedBody)) {
		return fmt.Errorf("unexpected response body: %s. Expected: %s", string(respBody), string(expectedBody))
	}

	return nil
}
