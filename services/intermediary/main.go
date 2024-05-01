package main

import (
	"modular-monolith-boilerplate/pkg/restapi"
	"modular-monolith-boilerplate/services/intermediary/adapter/controller"
	_ "modular-monolith-boilerplate/services/intermediary/adapter/repository/micro"
)

//nolint:unused
func main() {
	controller.RegisterRouting()
	_ = restapi.Run("8080", "intermediary")
}
