package restapi

import (
	"encoding/json"
	"io"
	"modular-monolith-boilerplate/pkg/errors"
	"modular-monolith-boilerplate/pkg/logger"
	"net/http"
	"os"
	"strconv"
)

var InternalApiBaseURL string

func init() {
	path := os.Getenv("INTERNAL_API_BASE_URL")
	if path == "" {
		InternalApiBaseURL = "http://localhost:8080"
	}
}

type RestClient struct {
}

func NewRestClient() *RestClient {
	return &RestClient{}
}

func (rc *RestClient) CallGet(path string, responseType interface{}) (interface{}, *errors.ApiError) {
	resp, err := http.Get(path)
	// URLがnilだったり、Timeoutが発生した場合といったサーバーとの疎通前のエラーを検証する。
	if err != nil {
		logger.Default().Error("Error Request:", err)
		return nil, errors.NewSystemError(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, &errors.ApiError{
			Code:    resp.StatusCode,
			Message: "calling API returned " + strconv.Itoa(resp.StatusCode) + " status",
		}
	}

	body, _ := io.ReadAll(resp.Body)
	_, ok := responseType.(string)
	if ok {
		var message string
		err := json.Unmarshal(body, &message)
		if err != nil {
			logger.Default().Error("Error Unmarshaling JSON:", err)
			return nil, errors.NewSystemError(err)
		}
		return message, nil
	}
	err = json.Unmarshal(body, responseType)
	if err != nil {
		logger.Default().Error("Error Unmarshaling JSON:", err)
		return nil, errors.NewSystemError(err)
	}
	return responseType, nil

}
