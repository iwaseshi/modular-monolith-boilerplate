package micro

import (
	"fmt"
	"modular-monolith-boilerplate/pkg/di"
	"modular-monolith-boilerplate/pkg/errors"
	"modular-monolith-boilerplate/pkg/restapi"
	"modular-monolith-boilerplate/services/intermediary/domain/repository"
)

var (
	healthCheckBasePath = "http://localhost:8080"
)

func init() {
	di.RegisterBean(NewMicroHealthCheckRepository)
	if restapi.IsRunningOnCloud() {
		healthCheckBasePath = "https://healthcheck-rftndbrsdq-an.a.run.app"
	}
}

type MicroHealthCheckRepository struct {
	restClient *restapi.RestClient
}

func NewMicroHealthCheckRepository() repository.HealthCheckRepository {
	return &MicroHealthCheckRepository{
		restClient: restapi.NewRestClient(),
	}
}

func (hcr *MicroHealthCheckRepository) Ping(c *restapi.Context) (*string, *errors.ApiError) {
	message := ""
	resp, err := hcr.restClient.CallGet(healthCheckBasePath+"/health-check/ping", message)
	if err != nil {
		return nil, err
	}
	message, ok := resp.(string)
	if !ok {
		return nil, errors.NewSystemError(fmt.Errorf("unexpected response type"))
	}
	message = message + " from rest call"
	return &message, nil
}
