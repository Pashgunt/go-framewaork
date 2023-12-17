package contract

import "net/http"

type RequestContract interface {
	GetRequest() http.Request
}
