package usecase

import (
	"context"
	"modular-monolith-boilerplate/services/healthcheck/domain"

	"modular-monolith-boilerplate/pkg/di"
	"modular-monolith-boilerplate/pkg/errors"
)

func init() {
	di.RegisterBean(NewHealthCheckInteractor)
}

type HealthCheckUseCase interface {
	Ping(c context.Context) (*string, *errors.ApiError)
	Readiness(c context.Context, req *domain.ReadyRequest) (*domain.ReadyResponse, *errors.ApiError)
}

type HealthCheckInteractor struct {
}

func NewHealthCheckInteractor() HealthCheckUseCase {
	return &HealthCheckInteractor{}

}

func (hci *HealthCheckInteractor) Ping(c context.Context) (*string, *errors.ApiError) {
	message := "pong"
	return &message, nil
}

func (hci *HealthCheckInteractor) Readiness(c context.Context, req *domain.ReadyRequest) (*domain.ReadyResponse, *errors.ApiError) {
	var res = domain.ReadyResponse{
		Message: "yeah!!",
	}
	if req.Shout != "Are you ready?" {
		res.Message = "no!"
		return &res, nil
	}
	return &res, nil
}
