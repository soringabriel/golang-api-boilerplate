package endpoints

import (
	"api/app/middlewares"
	"net/http"
)

type Endpoint struct {
	Path        string
	Middlewares []middlewares.Middleware
	HandlerFunc http.HandlerFunc
}

func (e *Endpoint) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler := e.HandlerFunc
	for _, middleware := range e.Middlewares {
		handler = middleware(handler)
	}
	handler(w, r)
}
