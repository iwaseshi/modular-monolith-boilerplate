package mono

import (
	"modular-monolith-boilerplate/pkg/di"
	"modular-monolith-boilerplate/pkg/restapi"
	"modular-monolith-boilerplate/services/healthcheck/usecase"
	"modular-monolith-boilerplate/services/intermediary/domain/repository"
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

func (hcr *MonoHealthCheckRepository) Ping(c *restapi.Context) (string, error) {
	message, err := hcr.healthCheckUseCase.Ping(c)
	if err != nil {
		return "", err
	}
	return *message + " from function", nil
}
