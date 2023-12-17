package app

import (
	"log"
	"net/http"
)

type API struct {
	router *http.ServeMux
}

func NewApi() *API {
	return &API{
		router: http.NewServeMux(),
	}
}

func (api API) Init(port string) {
	log.Fatal(http.ListenAndServe(port, api.router))
}

func (api API) GET(path string, callable func(writer http.ResponseWriter, request *http.Request)) {
	api.router.HandleFunc(path, callable)
}
