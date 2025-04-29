package middlewares

import (
	"net/http"
	"strings"

	"api/app/responses/bad_request_responses"

	"api/helpers"
)

func AuthMiddlewareFactory() Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
				bad_request_responses.MissingAuthResponse.Write(w)
				return
			}

			token := strings.TrimPrefix(authHeader, "Bearer ")
			if token != helpers.GetEnvVariable("AUTH_TOKEN") {
				bad_request_responses.WrongAuthResponse.Write(w)
				return
			}

			next(w, r)
		}
	}
}

func ApiKeyMiddlewareFactory() Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			apiKey := r.URL.Query().Get("api_key")
			if apiKey == "" {
				bad_request_responses.MissingAuthResponse.Write(w)
				return
			}

			if apiKey != helpers.GetEnvVariable("AUTH_API_KEY") {
				bad_request_responses.WrongAuthResponse.Write(w)
				return
			}

			next(w, r)
		}
	}
}
