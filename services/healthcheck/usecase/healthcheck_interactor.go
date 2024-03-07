package usecase

import (
	"modular-monolith-boilerplate/services/healthcheck/domain"

	"modular-monolith-boilerplate/pkg/di"
	"modular-monolith-boilerplate/pkg/restapi"
)

func init() {
	di.RegisterBean(NewHealthCheckInteractor)
}

type HealthCheckUseCase interface {
	Ping(c *restapi.Context) (*string, error)
	Readiness(c *restapi.Context, req *domain.ReadyRequest) (*domain.ReadyResponse, error)
}

type HealthCheckInteractor struct {
}

func NewHealthCheckInteractor() HealthCheckUseCase {
	return &HealthCheckInteractor{}

}

func (hci *HealthCheckInteractor) Ping(c *restapi.Context) (*string, error) {
	message := "pong"
	return &message, nil
}

func (hci *HealthCheckInteractor) Readiness(c *restapi.Context, req *domain.ReadyRequest) (*domain.ReadyResponse, error) {
	var res = domain.ReadyResponse{
		Message: "yeah!!",
	}
	if req.Shout != "Are you ready?" {
		res.Message = "no!"
		return &res, nil
	}
	return &res, nil
}
