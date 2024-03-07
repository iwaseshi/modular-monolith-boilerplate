package main

import (
	"modular-monolith-boilerplate/pkg/restapi"
	"os"

	intermediary "modular-monolith-boilerplate/services/intermediary/adapter/controller"

	// デフォルトではmonoModeで起動する。microで起動する場合は以下のコメントを外しmonoをコメントアウトする。
	//_ "modular-monolith-boilerplate/services/intermediary/adapter/repository/micro"
	healthcheck "modular-monolith-boilerplate/services/healthcheck/adapter/controller"
	_ "modular-monolith-boilerplate/services/intermediary/adapter/repository/mono"
)

func main() {
	healthcheck.RegisterRouting()
	intermediary.RegisterRouting()
	port := os.Getenv("PORT")
	if port == "" {
		port = restapi.DefaultPort
	}
	_ = restapi.Run(port)
}
