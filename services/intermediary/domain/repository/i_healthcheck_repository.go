package repository

import "modular-monolith-boilerplate/pkg/restapi"

type HealthCheckRepository interface {
	Ping(c *restapi.Context) (string, error)
}
