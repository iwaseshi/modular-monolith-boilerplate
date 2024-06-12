package controller

import (
	"modular-monolith-boilerplate/pkg/adapter"
	"modular-monolith-boilerplate/pkg/di"
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
	storage adapter.Storage
}

func NewFileServiceController(storage adapter.Storage) *FileServiceController {
	return &FileServiceController{
		storage: storage,
	}
}

func (fsc *FileServiceController) Upload(c *restapi.Context) {

	file, header, err := c.FormFile("file")
	if err != nil {
		c.ApiResponse(http.StatusBadRequest, err.Error())
		return
	}
	defer file.Close()

	bucket, err := fsc.storage.NewBucketConnection(c.Context(), "fileservice-app-bucket")
	if err != nil {
		c.ApiResponse(http.StatusInternalServerError, err.Error())
		return
	}
	defer bucket.Close()

	bucket.AddFile(file, header.Filename, header.Header.Get("Content-Type"))
	err = bucket.WriteFiles()
	if err != nil {
		c.ApiResponse(http.StatusInternalServerError, err.Error())
		return
	}
	c.ApiResponse(http.StatusOK, "File uploaded successfully")
}
