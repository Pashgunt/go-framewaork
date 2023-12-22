package decorator

import (
	"net/http"
	"project/dto"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)
type DispatcherFunc func(data *dto.ApiRequestDto)
type AppHandleDecorator func(pattern string, handler func(http.ResponseWriter, *http.Request))
type Return func(pattern string, handler HandlerFunc, dispatcher DispatcherFunc, data *dto.ApiRequestDto)

func App(fn AppHandleDecorator) Return {
	return func(pattern string, handler HandlerFunc, dispatcher DispatcherFunc, data *dto.ApiRequestDto) {
		fn(pattern, handler)
		dispatcher(data)
	}
}
