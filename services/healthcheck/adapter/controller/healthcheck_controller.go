package controller

import (
	"fmt"
	"modular-monolith-boilerplate/pkg/di"
	"modular-monolith-boilerplate/pkg/restapi"
	"modular-monolith-boilerplate/services/healthcheck/domain"
	"modular-monolith-boilerplate/services/healthcheck/usecase"
)

func init() {
	di.RegisterBean(NewHealthCheckController)
}

func RegisterRouting() {
	_ = di.GetContainer().Invoke(
		func(hcc *HealthCheckController) {
			group := restapi.NewGroup("/health-check")
			group.RegisterGET("/ping", restapi.Handler(hcc.Ping))
			group.RegisterPOST("/readiness", restapi.Handler(hcc.Readiness))
		},
	)
}

type HealthCheckController struct {
	healthCheckUseCase usecase.HealthCheckUseCase
}

func NewHealthCheckController(healthCheckUseCase usecase.HealthCheckUseCase) *HealthCheckController {
	return &HealthCheckController{
		healthCheckUseCase: healthCheckUseCase,
	}
}

func (hcc *HealthCheckController) Ping(c *restapi.Context) {
	message, err := hcc.healthCheckUseCase.Ping(c)
	if err != nil {
		c.GinContext().JSON(500, err)
		return
	}
	c.GinContext().JSON(200, message)
}

func (hcc *HealthCheckController) Readiness(c *restapi.Context) {
	req := &domain.ReadyRequest{}
	if err := c.GinContext().BindJSON(req); err != nil {
		fmt.Println(err.Error())
		c.GinContext().JSON(400, err)
		return
	}
	res, err := hcc.healthCheckUseCase.Readiness(c, req)
	if err != nil {
		c.GinContext().JSON(500, err.Error())
		return
	}
	c.GinContext().JSON(200, res)
}
