package controller

import (
	"net/http"
	"project/helper"
)

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}

func (u UserController) GetPingInfo(writer http.ResponseWriter, request *http.Request) {
	pipeline := helper.Pipeline{}
	pipeline.SetRequestData(helper.NewHandle(writer, *request).SetMethod(http.MethodGet).Validate(
		map[string]helper.ValidateRule{
			"user_name": {Required: false, Regex: "."},
		},
	))
}
