package usecase

import (
	"modular-monolith-boilerplate/pkg/di"
	errors "modular-monolith-boilerplate/pkg/dto"
	"modular-monolith-boilerplate/services/callanotherapi/domain/repository"

	"github.com/gin-gonic/gin"
)

func init() {
	di.RegisterBean(NewCallAnotherApiInteractorr)
}

type CallAnotherApiUseCase interface {
	Call(c *gin.Context) (*string, *errors.ApiError)
}

type CallAnotherApiInteractor struct {
	healthCheckRepository repository.HealthCheckRepository
}

func NewCallAnotherApiInteractorr(healthCheckRepository repository.HealthCheckRepository) CallAnotherApiUseCase {
	return &CallAnotherApiInteractor{
		healthCheckRepository: healthCheckRepository,
	}

}

func (hci *CallAnotherApiInteractor) Call(c *gin.Context) (*string, *errors.ApiError) {
	message, err := hci.healthCheckRepository.Ping(c)
	if err != nil {
		return nil, errors.NewSystemError(err)
	}
	return &message, nil
}
