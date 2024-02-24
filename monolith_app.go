package main

import (
	"modular-monolith-boilerplate/pkg/restapi"

	callanotherapi "modular-monolith-boilerplate/services/callanotherapi/adapter/controller"

	// デフォルトではmonoModeで起動する。microで起動する場合は以下のコメントを外しmonoをコメントアウトする。
	//_ "modular-monolith-boilerplate/services/callanotherapi/adapter/repository/micro"
	_ "modular-monolith-boilerplate/services/callanotherapi/adapter/repository/mono"
	healthcheck "modular-monolith-boilerplate/services/healthcheck/adapter/controller"
)

func main() {
	healthcheck.RegisterRouting()
	callanotherapi.RegisterRouting()
	_ = restapi.Run(restapi.DefaultPort)
}
