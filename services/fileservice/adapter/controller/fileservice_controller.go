package controller

import (
	"modular-monolith-boilerplate/pkg/di"
	"modular-monolith-boilerplate/pkg/gcp"
	"modular-monolith-boilerplate/pkg/restapi"
	"net/http"
)

func init() {
	di.RegisterBean(NewFileServiceController)
}

func RegisterRouting() {
	_ = di.GetContainer().Invoke(
		func(fsc *FileServiceController) {
			group := restapi.NewGroup("/file-service")
			group.RegisterPOST("/upload", fsc.Upload)
		},
	)
}

type FileServiceController struct {
}

func NewFileServiceController() *FileServiceController {
	return &FileServiceController{}
}

func (fsc *FileServiceController) Upload(c *restapi.Context) {

	file, header, err := c.FormFile("file")
	if err != nil {
		c.ApiResponse(http.StatusBadRequest, err.Error())
		return
	}
	defer file.Close()

	storage, err := gcp.NewCloudStorage(c.Context(), "modular-monolith-sample-app-bucket")
	if err != nil {
		c.ApiResponse(http.StatusInternalServerError, err.Error())
		return
	}
	defer storage.Close()

	storage.AddFile(file, header.Filename, header.Header.Get("Content-Type"))
	err = storage.WriteFiles()
	if err != nil {
		c.ApiResponse(http.StatusInternalServerError, err.Error())
		return
	}
	c.ApiResponse(http.StatusOK, "File uploaded successfully")
}
