package requests

import "net/http"

type Request interface {
	Defaults()
	Read(r *http.Request)
	Validate() error
	GetMethod() string
	GetParams() []byte
}
