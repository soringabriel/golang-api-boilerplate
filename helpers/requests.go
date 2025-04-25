package helpers

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type HttpRequestParams struct {
	Method      string
	Url         string
	QueryParams map[string]string
	Headers     map[string]string
	RequestBody []byte
}

func HttpRequest(params HttpRequestParams) (*http.Response, string, error) {
	client := &http.Client{}

	// Build the url
	requestUrl, err := url.Parse(params.Url)
	if err != nil {
		return nil, "", fmt.Errorf("[HttpRequest] Error parsing url: %v", err)
	}
	queryParams := url.Values{}
	for key, value := range params.QueryParams {
		queryParams.Add(key, value)
	}
	requestUrl.RawQuery = queryParams.Encode()

	// Build the request
	req, err := http.NewRequest(params.Method, requestUrl.String(), bytes.NewBuffer(params.RequestBody))
	if err != nil {
		return nil, "", fmt.Errorf("[HttpRequest] Error creating request: %v", err)
	}
	for key, value := range params.Headers {
		req.Header.Add(key, value)
	}

	// Make request
	resp, err := client.Do(req)
	if err != nil {
		return nil, "", fmt.Errorf("[HttpRequest] Error making request: %v", err)
	}
	defer resp.Body.Close()

	// Get response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, "", fmt.Errorf("[HttpRequest] Error reading response body: %v", err)
	}

	return resp, string(body), nil
}
