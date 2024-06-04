package restapi

import (
	"context"
	"mime/multipart"
	"modular-monolith-boilerplate/pkg/errors"
	"modular-monolith-boilerplate/pkg/logger"
	"os"

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

func (rg *routerGroup) RegisterGET(getPath string, fun HandlerFunc) {
	rg.group.GET(getPath, convertHandler(fun))
}

func (rg *routerGroup) RegisterPOST(postPath string, fun HandlerFunc) {
	rg.group.POST(postPath, convertHandler(fun))
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

func (c *Context) Context() context.Context {
	return c.stdCtx
}

func (c *Context) BindJson(req any) (error *errors.ApiError) {
	if err := c.ginCtx.BindJSON(req); err != nil {
		logger.WithCtx(c.stdCtx).Error("Error binding request: %s", err.Error())
		return errors.NewBusinessError(err)
	}
	return nil
}

func (c *Context) FormFile(key string) (multipart.File, *multipart.FileHeader, error) {
	return c.ginCtx.Request.FormFile(key)
}

func (c *Context) ApiResponse(statusCode int, body interface{}) {
	c.ginCtx.JSON(statusCode, body)
}

type HandlerFunc func(*Context)

func convertHandler(handler HandlerFunc) gin.HandlerFunc {
	return func(ginCtx *gin.Context) {
		customCtx := NewContext(ginCtx)
		logger.RegisterInCtx(customCtx.Context())
		handler(customCtx)
	}
}

func IsRunningOnCloud() bool {
	// Cloud Runは`K_SERVICE`環境変数を提供している
	_, exists := os.LookupEnv("K_SERVICE")
	return exists
}
