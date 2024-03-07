package controller

import (
	"modular-monolith-boilerplate/pkg/di"
	"modular-monolith-boilerplate/pkg/restapi"
	"modular-monolith-boilerplate/services/intermediary/usecase"
)

func init() {
	di.RegisterBean(NewIntermediaryController)
}

func RegisterRouting() {
	_ = di.GetContainer().Invoke(
		func(ic *IntermediaryController) {
			group := restapi.NewGroup("/call-another-api")
			group.RegisterGET("/call", restapi.Handler(ic.Call))
		},
	)
}

type IntermediaryController struct {
	intermediaryUseCase usecase.IntermediaryUseCase
}

func NewIntermediaryController(intermediaryUseCase usecase.IntermediaryUseCase) *IntermediaryController {
	return &IntermediaryController{
		intermediaryUseCase: intermediaryUseCase,
	}
}

func (ic *IntermediaryController) Call(c *restapi.Context) {
	message, err := ic.intermediaryUseCase.Call(c)
	if err != nil {
		c.ApiResponse(err.Code, err.Error())
		return
	}
	c.ApiResponse(200, message)
}
