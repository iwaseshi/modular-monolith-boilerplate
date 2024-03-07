package main

import (
	"modular-monolith-boilerplate/pkg/restapi"
	"modular-monolith-boilerplate/services/callanotherapi/adapter/controller"
	_ "modular-monolith-boilerplate/services/callanotherapi/adapter/repository/micro"
)

//nolint:unused
func main() {
	controller.RegisterRouting()
	_ = restapi.Run("8080")
}
