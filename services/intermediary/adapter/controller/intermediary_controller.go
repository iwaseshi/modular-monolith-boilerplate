package controller

import (
	"modular-monolith-boilerplate/pkg/di"
	"modular-monolith-boilerplate/pkg/restapi"
	"modular-monolith-boilerplate/services/intermediary/usecase"

	"github.com/gin-gonic/gin"
)

func init() {
	di.RegisterBean(NewIntermediaryController)
}

func RegisterRouting() {
	_ = di.GetContainer().Invoke(
		func(caa *IntermediaryApiController) {
			group := restapi.NewGroup("/call-another-api")
			group.RegisterGET("/call", caa.Call)
		},
	)
}

type IntermediaryApiController struct {
	intermediaryUseCase usecase.IntermediaryUseCase
}

func NewIntermediaryController(intermediaryUseCase usecase.IntermediaryUseCase) *IntermediaryApiController {
	return &IntermediaryApiController{
		intermediaryUseCase: intermediaryUseCase,
	}
}

func (caa *IntermediaryApiController) Call(c *gin.Context) {
	message, err := caa.intermediaryUseCase.Call(c)
	if err != nil {
		c.JSON(err.Code, err.Message)
		return
	}
	c.JSON(200, message)
}
