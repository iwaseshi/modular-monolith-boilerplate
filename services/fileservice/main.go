package main

import (
	_ "modular-monolith-boilerplate/pkg/gcp"
	"modular-monolith-boilerplate/pkg/restapi"
	"modular-monolith-boilerplate/services/fileservice/adapter/controller"
)

//nolint:unused
func main() {
	controller.RegisterRouting()
	_ = restapi.Run(restapi.DefaultPort)
}
