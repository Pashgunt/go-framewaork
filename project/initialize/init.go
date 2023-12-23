package initialize

import (
	"log"
	"net/http"
	"net/url"
	"project/decorator"
	"project/dto"
	"project/enum"
	"project/event"
	"project/helper"
	"sync"
)

type API struct {
	router           *http.ServeMux
	publisher        event.Publisher
	requestValidator helper.ValidateStructureRequest
}

func NewApi() *API {
	return &API{
		router:           http.NewServeMux(),
		publisher:        event.Publisher{},
		requestValidator: helper.ValidateStructureRequest{},
	}
}

func (api API) Init(port string) {
	api.initErrorProcessing()
	api.initFile()
	log.Fatal(http.ListenAndServe(port, api.router))
}

// TODO подумать как перехватить глобальный объект http.Request в этом месте для всех методов
func (api API) GET(path string, callable func(writer http.ResponseWriter, request *http.Request)) {
	decorator.App(api.router.HandleFunc)(path, callable, api.eventDispatcher, dto.NewApiRequestDto(http.Request{
		Method: "GET",
		URL: &url.URL{
			Path: path,
		},
	}, "GET", path))
}

func (api API) POST(path string, callable func(writer http.ResponseWriter, request *http.Request)) {
	decorator.App(api.router.HandleFunc)(path, callable, api.eventDispatcher, dto.NewApiRequestDto(http.Request{
		Method: "GET",
		URL: &url.URL{
			Path: path,
		},
	}, "GET", path))
}

func (api API) eventDispatcher(data *dto.ApiRequestDto) {
	var wg sync.WaitGroup
	mux := sync.Mutex{}
	subscriber := api.publisher.RequestSubscribe()
	wg.Add(1)
	api.publisher.RequestPublish(data)
	go func() {
		defer wg.Done()
		mux.Lock()
		api.requestValidator.Validate(<-subscriber)
		mux.Unlock()
	}()
	wg.Wait()
	close(subscriber)
}

func (api API) initErrorProcessing() {
	defer func() {
		if err := recover(); err != nil {
			log.Fatalf(enum.SystemMessageError, "001", err)
		}
	}()
}

func (api API) initFile() {
	file := file{}
	file.Init()
}
