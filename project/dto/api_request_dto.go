package dto

import "net/http"

type ApiRequestDto struct {
	request     http.Request
	innerMethod string
	innerPath   string
}

func NewApiRequestDto(request http.Request, innerMethod string, innerPath string) *ApiRequestDto {
	return &ApiRequestDto{request: request, innerMethod: innerMethod, innerPath: innerPath}
}

func (ard ApiRequestDto) Request() http.Request {
	return ard.request
}

func (ard ApiRequestDto) InnerMethod() string {
	return ard.innerMethod
}

func (ard ApiRequestDto) InnerPath() string {
	return ard.innerPath
}
