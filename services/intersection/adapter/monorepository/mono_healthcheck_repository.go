package repository

import (
	"modular-monolith-boilerplate/pkg/di"
	"modular-monolith-boilerplate/services/callanotherapi/domain/repository"
	"modular-monolith-boilerplate/services/healthcheck/usecase"

	"github.com/gin-gonic/gin"
)

func init() {
	di.RegisterBean(NewMonoHealthCheckRepository)
}

type MonoHealthCheckRepository struct {
	healthCheckUseCase usecase.HealthCheckUseCase
}

func NewMonoHealthCheckRepository(healthCheckUseCase usecase.HealthCheckUseCase) repository.HealthCheckRepository {
	return &MonoHealthCheckRepository{
		healthCheckUseCase: healthCheckUseCase,
	}
}

func (hcr *MonoHealthCheckRepository) Ping(c *gin.Context) (string, error) {
	message, err := hcr.healthCheckUseCase.Ping(c)
	if err != nil {
		return "", err
	}
	return *message + " from function", nil
}
