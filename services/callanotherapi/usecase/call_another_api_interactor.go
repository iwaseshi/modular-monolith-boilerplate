package usecase

import (
	"modular-monolith-boilerplate/pkg/di"
	"modular-monolith-boilerplate/services/callanotherapi/domain/repository"

	"github.com/gin-gonic/gin"
)

func init() {
	di.RegisterBean(NewCallAnotherApiInteractorr)
}

type CallAnotherApiUseCase interface {
	Call(c *gin.Context) (*string, error)
}

type CallAnotherApiInteractor struct {
	healthCheckRepository repository.HealthCheckRepository
}

func NewCallAnotherApiInteractorr(healthCheckRepository repository.HealthCheckRepository) CallAnotherApiUseCase {
	return &CallAnotherApiInteractor{
		healthCheckRepository: healthCheckRepository,
	}

}

func (hci *CallAnotherApiInteractor) Call(c *gin.Context) (*string, error) {
	message, err := hci.healthCheckRepository.Ping(c)
	if err != nil {
		return nil, err
	}
	return &message, nil
}
