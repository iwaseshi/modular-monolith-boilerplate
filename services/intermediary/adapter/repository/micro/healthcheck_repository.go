package micro

import (
	"fmt"
	"modular-monolith-boilerplate/pkg/config"
	"modular-monolith-boilerplate/pkg/di"
	"modular-monolith-boilerplate/pkg/errors"
	"modular-monolith-boilerplate/pkg/restapi"
	"modular-monolith-boilerplate/services/intermediary/domain/repository"
	"os"
)

func init() {
	di.RegisterBean(NewMicroHealthCheckRepository)
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
	mode := os.Getenv("MODE")
	if mode == "" {
		mode = "mono"
	}
	healthCheckBasePath := config.Get("healthCheckBasePath." + mode)
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
