package usecase

import (
	"modular-monolith-boilerplate/pkg/di"
	"modular-monolith-boilerplate/pkg/errors"
	"modular-monolith-boilerplate/pkg/logger"
	"modular-monolith-boilerplate/pkg/restapi"
	"modular-monolith-boilerplate/services/intermediary/domain/repository"
)

func init() {
	di.RegisterBean(NewIntermediaryInteractorr)
}

type IntermediaryUseCase interface {
	Call(c *restapi.Context) (*string, *errors.ApiError)
}

type IntermediaryInteractor struct {
	healthCheckRepository repository.HealthCheckRepository
}

func NewIntermediaryInteractorr(healthCheckRepository repository.HealthCheckRepository) IntermediaryUseCase {
	return &IntermediaryInteractor{
		healthCheckRepository: healthCheckRepository,
	}

}

func (ii *IntermediaryInteractor) Call(c *restapi.Context) (*string, *errors.ApiError) {
	message, err := ii.healthCheckRepository.Ping(c)
	if err != nil {
		logger.WithCtx(c.Context()).Error("Failed to Call Health Check API", err.Unwrap())
		return nil, errors.NewSystemError(err.Unwrap())
	}
	return message, nil
}
