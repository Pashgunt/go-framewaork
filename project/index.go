package main

import (
	"project/controller"
	"project/initialize"
)

func main() {
	api := initialize.NewApi()

	api.GET("/", (controller.NewUserController()).GetPingInfo)
	api.Init(":8081")
}
