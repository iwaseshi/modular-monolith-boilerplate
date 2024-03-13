package main

import (
	"modular-monolith-boilerplate/pkg/config"
	"modular-monolith-boilerplate/pkg/restapi"

	intermediary "modular-monolith-boilerplate/services/intermediary/adapter/controller"

	// デフォルトではmonoModeで起動する。microで起動する場合は以下のコメントを外しmonoをコメントアウトする。
	healthcheck "modular-monolith-boilerplate/services/healthcheck/adapter/controller"
	_ "modular-monolith-boilerplate/services/intermediary/adapter/repository/micro"
	//_ "modular-monolith-boilerplate/services/intermediary/adapter/repository/mono"
)

func main() {
	config.LoadServiceConfig(".")
	healthcheck.RegisterRouting()
	intermediary.RegisterRouting()
	_ = restapi.Run("8080")
}
