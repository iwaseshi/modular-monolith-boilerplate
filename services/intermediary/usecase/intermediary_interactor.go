package usecase

import (
	"modular-monolith-boilerplate/pkg/di"
	"modular-monolith-boilerplate/pkg/errors"
	"modular-monolith-boilerplate/services/intermediary/domain/repository"

	"github.com/gin-gonic/gin"
)

func init() {
	di.RegisterBean(NewIntermediaryInteractorr)
}

type IntermediaryUseCase interface {
	Call(c *gin.Context) (*string, *errors.ApiError)
}

type IntermediaryInteractor struct {
	healthCheckRepository repository.HealthCheckRepository
}

func NewIntermediaryInteractorr(healthCheckRepository repository.HealthCheckRepository) IntermediaryUseCase {
	return &IntermediaryInteractor{
		healthCheckRepository: healthCheckRepository,
	}

}

func (hci *IntermediaryInteractor) Call(c *gin.Context) (*string, *errors.ApiError) {
	message, err := hci.healthCheckRepository.Ping(c)
	if err != nil {
		return nil, errors.NewSystemError(err)
	}
	return &message, nil
}
