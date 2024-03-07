package main

import (
	"modular-monolith-boilerplate/pkg/restapi"
	"os"

	callanotherapi "modular-monolith-boilerplate/services/callanotherapi/adapter/controller"

	// デフォルトではmonoModeで起動する。microで起動する場合は以下のコメントを外しmonoをコメントアウトする。
	//_ "modular-monolith-boilerplate/services/callanotherapi/adapter/repository/micro"
	_ "modular-monolith-boilerplate/services/callanotherapi/adapter/repository/mono"
	healthcheck "modular-monolith-boilerplate/services/healthcheck/adapter/controller"
)

func main() {
	healthcheck.RegisterRouting()
	callanotherapi.RegisterRouting()
	port := os.Getenv("PORT")
	if port == "" {
		port = restapi.DefaultPort
	}
	_ = restapi.Run(port)
}
