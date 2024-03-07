package restapi

import (
	"github.com/gin-gonic/gin"
)

const (
	DefaultPort = "8080"
)

var (
	router = gin.Default()
)

type routerGroup struct {
	path  string
	group *gin.RouterGroup
}

func NewGroup(groupPath string) *routerGroup {
	return &routerGroup{
		groupPath,
		router.Group(groupPath),
	}
}

func (rg *routerGroup) RegisterGET(getPath string, fun gin.HandlerFunc) {
	rg.group.GET(getPath, fun)
}

func (rg *routerGroup) RegisterPOST(postPath string, fun gin.HandlerFunc) {
	rg.group.POST(postPath, fun)
}

func Run(port string) error {
	if port == "" {
		port = DefaultPort
	}
	return router.Run(":" + port)
}
