package contract

import "net/http"

type ResponseContract interface {
	GetResponse() http.ResponseWriter
}
