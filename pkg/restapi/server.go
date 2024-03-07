package restapi

import (
	"context"
	"modular-monolith-boilerplate/pkg/logger"

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

type Context struct {
	ginCtx *gin.Context
	stdCtx context.Context
}

func NewContext(ginCtx *gin.Context) *Context {
	return &Context{
		ginCtx: ginCtx,
		stdCtx: ginCtx.Request.Context(),
	}
}

func (c *Context) GinContext() *gin.Context {
	return c.ginCtx
}

func (c *Context) StandardContext() context.Context {
	return c.stdCtx
}

func (c *Context) ApiResponse(statusCode int, body interface{}) {
	c.ginCtx.JSON(statusCode, body)
}

type HandlerFunc func(*Context)

func Handler(handler HandlerFunc) gin.HandlerFunc {
	return func(ginCtx *gin.Context) {
		customCtx := NewContext(ginCtx)
		logger.RegisterInCtx(customCtx.StandardContext())
		handler(customCtx)
	}
}
