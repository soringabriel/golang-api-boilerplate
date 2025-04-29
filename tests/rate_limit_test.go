package tests

import (
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"
	"testing"

	"api/app/endpoints/user_endpoints"
	"api/app/responses/bad_request_responses"
	"api/helpers"
)

func TestGeneralApiRateLimit(t *testing.T) {
	t.Parallel()

	responses429 := int64(0)
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			resp, _, err := helpers.HttpRequest(helpers.HttpRequestParams{
				Method: http.MethodGet,
				Url: fmt.Sprintf(
					"http://%s%s",
					helpers.GetEnvVariable("API_IP_PORT"),
					user_endpoints.ReadUserEndpoint.Path,
				),
				Headers: map[string]string{
					"Accept":        "application/json",
					"Authorization": fmt.Sprintf("Bearer %s", helpers.GetEnvVariable("AUTH_TOKEN")),
				},
			})
			if err != nil {
				return
			}
			if resp.StatusCode == bad_request_responses.RateLimitResponse.StatusCode {
				atomic.AddInt64(&responses429, 1)
			}
		}()
	}
	wg.Wait()

	if responses429 != 5 {
		t.Fatalf("Expected 5 429 responses. Got %d", responses429)
	}
}

func TestEndpointIPRateLimit(t *testing.T) {
	t.Parallel()

	responses429 := int64(0)
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			resp, _, err := helpers.HttpRequest(helpers.HttpRequestParams{
				Method: http.MethodGet,
				Url: fmt.Sprintf(
					"http://%s%s",
					helpers.GetEnvVariable("API_IP_PORT"),
					user_endpoints.ReadLimitedUserEndpoint.Path,
				),
				Headers: map[string]string{
					"Accept":        "application/json",
					"Authorization": fmt.Sprintf("Bearer %s", helpers.GetEnvVariable("AUTH_TOKEN")),
				},
			})
			if err != nil {
				return
			}
			if resp.StatusCode == bad_request_responses.RateLimitResponse.StatusCode {
				atomic.AddInt64(&responses429, 1)
			}
		}()
	}
	wg.Wait()

	if responses429 != 9 {
		t.Fatalf("Expected 9 429 responses. Got %d", responses429)
	}
}
