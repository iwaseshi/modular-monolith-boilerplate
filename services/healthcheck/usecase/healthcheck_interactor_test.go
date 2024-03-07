package usecase_test

import (
	"modular-monolith-boilerplate/pkg/restapi"
	"modular-monolith-boilerplate/services/healthcheck/domain"
	"modular-monolith-boilerplate/services/healthcheck/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthCheckInteractor_Ping(t *testing.T) {
	var ctx *restapi.Context

	hci := usecase.NewHealthCheckInteractor()

	res, err := hci.Ping(ctx)

	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, "pong", *res)
}

func TestHealthCheckInteractor_Readiness(t *testing.T) {
	var ctx *restapi.Context

	hci := usecase.NewHealthCheckInteractor()

	req := &domain.ReadyRequest{
		Shout: "Are you ready?",
	}
	res, err := hci.Readiness(ctx, req)

	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.IsType(t, &domain.ReadyResponse{}, res)
	assert.Equal(t, "yeah!!", res.Message)
}

func TestHealthCheckInteractor_Readiness_Shout(t *testing.T) {
	var ctx *restapi.Context
	hci := usecase.NewHealthCheckInteractor()

	req := &domain.ReadyRequest{
		Shout: "not ready",
	}
	res, err := hci.Readiness(ctx, req)

	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.IsType(t, &domain.ReadyResponse{}, res)
	assert.Equal(t, "no!", res.Message)
}
