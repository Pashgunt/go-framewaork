package enum

import "net/http"

const (
	get    = http.MethodGet
	post   = http.MethodPost
	put    = http.MethodPut
	patch  = http.MethodPatch
	delete = http.MethodDelete
)

func GetHttpMethodSlice() []string {
	return []string{get, post, put, patch, delete}
}
