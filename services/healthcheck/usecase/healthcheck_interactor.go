package usecase

import (
	"modular-monolith-boilerplate/services/healthcheck/domain"

	"modular-monolith-boilerplate/pkg/di"

	"github.com/gin-gonic/gin"
)

func init() {
	di.RegisterBean(NewHealthCheckInteractor)
}

type HealthCheckUseCase interface {
	Ping(c *gin.Context) (*string, error)
	Readiness(c *gin.Context, req *domain.ReadyRequest) (*domain.ReadyResponse, error)
}

type HealthCheckInteractor struct {
}

func NewHealthCheckInteractor() HealthCheckUseCase {
	return &HealthCheckInteractor{}

}

func (hci *HealthCheckInteractor) Ping(c *gin.Context) (*string, error) {
	message := "pong"
	return &message, nil
}

func (hci *HealthCheckInteractor) Readiness(c *gin.Context, req *domain.ReadyRequest) (*domain.ReadyResponse, error) {
	var res = domain.ReadyResponse{
		Message: "yeah!!",
	}
	if req.Shout != "Are you ready?" {
		res.Message = "no!"
		return &res, nil
	}
	return &res, nil
}
