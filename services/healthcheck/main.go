package main

import (
	"modular-monolith-boilerplate/pkg/restapi"
	"modular-monolith-boilerplate/services/healthcheck/adapter/controller"
)

//nolint:unused
func main() {
	controller.RegisterRouting()
	_ = restapi.Run(restapi.DefaultPort)
}
