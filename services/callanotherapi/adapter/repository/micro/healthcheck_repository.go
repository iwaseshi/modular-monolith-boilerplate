package micro

import (
	"fmt"
	"modular-monolith-boilerplate/pkg/di"
	"modular-monolith-boilerplate/pkg/restapi"
	"modular-monolith-boilerplate/services/callanotherapi/domain/repository"

	"github.com/gin-gonic/gin"
)

var (
	healthCheckBasePath = "https://healthcheck-rftndbrsdq-an.a.run.app/health-check/ping"
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

func (hcr *MicroHealthCheckRepository) Ping(c *gin.Context) (string, error) {
	message := ""
	resp, err := hcr.restClient.CallGet(healthCheckBasePath, message)
	if err != nil {
		return "", err
	}
	message, ok := resp.(string)
	if !ok {
		return "", fmt.Errorf("unexpected response type")
	}
	return message + " from rest call", nil
}
