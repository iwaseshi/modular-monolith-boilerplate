package main

import (
	"modular-monolith-boilerplate/pkg/di"
	"modular-monolith-boilerplate/pkg/restapi"

	ac "modular-monolith-boilerplate/services/callanotherapi/adapter/controller"
	hc "modular-monolith-boilerplate/services/healthcheck/adapter/controller"

	// micro mode
	_ "modular-monolith-boilerplate/services/intersection/adapter/microrepository"
	// mono mode
	//_ "modular-monolith-boilerplate/services/intersection/adapter/monorepository"
)

func main() {
	_ = di.GetContainer().Invoke(
		func(hcc *hc.HealthCheckController) {
			group := restapi.NewGroup("/health-check")
			group.RegisterGET("/ping", hcc.Ping)
			group.RegisterPOST("/readiness", hcc.Readiness)
		},
	)
	_ = di.GetContainer().Invoke(
		func(caa *ac.CallAnotherApiController) {
			group := restapi.NewGroup("/call-another-api")
			group.RegisterGET("/call", caa.Call)
		},
	)
	_ = restapi.Run()
}
