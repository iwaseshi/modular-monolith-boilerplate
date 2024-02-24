package controller

import (
	"modular-monolith-boilerplate/pkg/di"
	"modular-monolith-boilerplate/pkg/restapi"
	"modular-monolith-boilerplate/services/callanotherapi/usecase"

	"github.com/gin-gonic/gin"
)

func init() {
	di.RegisterBean(NewCallAnotherApiController)
}

func RegisterRouting() {
	_ = di.GetContainer().Invoke(
		func(caa *CallAnotherApiController) {
			group := restapi.NewGroup("/call-another-api")
			group.RegisterGET("/call", caa.Call)
		},
	)
}

type CallAnotherApiController struct {
	callAnotherApiUseCase usecase.CallAnotherApiUseCase
}

func NewCallAnotherApiController(callAnotherApiUseCase usecase.CallAnotherApiUseCase) *CallAnotherApiController {
	return &CallAnotherApiController{
		callAnotherApiUseCase: callAnotherApiUseCase,
	}
}

func (caa *CallAnotherApiController) Call(c *gin.Context) {
	message, err := caa.callAnotherApiUseCase.Call(c)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, message)
}
