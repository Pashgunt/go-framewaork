package main

import (
	"project/app"
	"project/controller"
)

func main() {
	api := app.NewApi()

	api.GET("/", (controller.NewUserController()).GetPingInfo)
	api.Init(":8081")
}
