package helpers

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"strings"
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

func GetRequestClientIP(r *http.Request) string {
	// Check X-Forwarded-For header
	ip := r.Header.Get("X-Forwarded-For")
	if ip != "" {
		// X-Forwarded-For may contain multiple IPs separated by commas
		parts := strings.Split(ip, ",")
		// Take the first non-empty IP and trim spaces
		for _, part := range parts {
			p := strings.TrimSpace(part)
			if p != "" {
				return p
			}
		}
	}

	// Fallback to X-Real-IP
	ip = r.Header.Get("X-Real-Ip")
	if ip != "" {
		return strings.TrimSpace(ip)
	}

	// Fallback to RemoteAddr
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return ip
}
