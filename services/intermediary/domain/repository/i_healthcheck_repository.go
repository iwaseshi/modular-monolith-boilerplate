package repository

import (
	"github.com/gin-gonic/gin"
)

type HealthCheckRepository interface {
	Ping(c *gin.Context) (string, error)
}
