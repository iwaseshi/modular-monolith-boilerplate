package main

import (
	"flag"
	"modular-monolith-boilerplate/pkg/grpc"
	"modular-monolith-boilerplate/pkg/restapi"
	"modular-monolith-boilerplate/services/healthcheck/adapter/controller"
	_ "modular-monolith-boilerplate/services/healthcheck/adapter/service"
)

var (
	runRpc = flag.Bool("rpc", true, "flag")
)

//nolint:unused
func main() {

	flag.Parse()
	// サンプル用。実際は使用するAPIのRun関数のみ記載すること
	if *runRpc {
		grpc.Run(grpc.DefaultPort)
	} else {
		controller.RegisterRouting()
		_ = restapi.Run(restapi.DefaultPort)
	}

}
