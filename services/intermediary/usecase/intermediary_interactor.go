package usecase

import (
	"modular-monolith-boilerplate/pkg/di"
	"modular-monolith-boilerplate/pkg/error"
	"modular-monolith-boilerplate/pkg/restapi"
	"modular-monolith-boilerplate/services/intermediary/domain/repository"
)

func init() {
	di.RegisterBean(NewIntermediaryInteractorr)
}

type IntermediaryUseCase interface {
	Call(c *restapi.Context) (*string, *error.ApiError)
}

type IntermediaryInteractor struct {
	healthCheckRepository repository.HealthCheckRepository
}

func NewIntermediaryInteractorr(healthCheckRepository repository.HealthCheckRepository) IntermediaryUseCase {
	return &IntermediaryInteractor{
		healthCheckRepository: healthCheckRepository,
	}

}

func (ii *IntermediaryInteractor) Call(c *restapi.Context) (*string, *error.ApiError) {
	message, err := ii.healthCheckRepository.Ping(c)
	if err != nil {
		return nil, error.NewSystemError(err)
	}
	return &message, nil
}
