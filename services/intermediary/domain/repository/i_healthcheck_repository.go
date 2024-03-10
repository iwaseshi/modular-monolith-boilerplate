package repository

import (
	"modular-monolith-boilerplate/pkg/errors"
	"modular-monolith-boilerplate/pkg/restapi"
)

type HealthCheckRepository interface {
	Ping(c *restapi.Context) (*string, *errors.ApiError)
}
