package controller_test

import (
	"bytes"
	"mime/multipart"
	"modular-monolith-boilerplate/pkg/adapter"
	"modular-monolith-boilerplate/pkg/restapi"
	"modular-monolith-boilerplate/services/fileservice/adapter/controller"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestFileServiceController_Upload(t *testing.T) {

	fsc := controller.NewFileServiceController(adapter.NewMockStorage())

	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	fw, err := w.CreateFormFile("file", "testfile.txt")
	assert.NoError(t, err)
	_, err = fw.Write([]byte("test content"))
	assert.NoError(t, err)
	w.Close()

	req := httptest.NewRequest("POST", "/file-service/upload", &b)
	req.Header.Set("Content-Type", w.FormDataContentType())
	rr := httptest.NewRecorder()
	ginCtx, _ := gin.CreateTestContext(rr)
	ginCtx.Request = req
	c := &restapi.Context{
		GinCtx: ginCtx,
		StdCtx: req.Context(),
	}

	fsc.Upload(c)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "\"File uploaded successfully\"", rr.Body.String())
}
