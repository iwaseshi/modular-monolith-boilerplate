package restapi

import (
	"encoding/json"
	"fmt"
	"io"
	"modular-monolith-boilerplate/pkg/logger"
	"net/http"
	"os"
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

func (rc *RestClient) CallGet(path string, responseType interface{}) (interface{}, error) {
	resp, err := http.Get(path)
	// URLがnilだったり、Timeoutが発生した場合といったサーバーとの疎通前のエラーを検証する。
	if err != nil {
		logger.Default().Error("Error Request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Error Response: %s", resp.Status)
	}

	body, _ := io.ReadAll(resp.Body)
	_, ok := responseType.(string)
	if ok {
		return string(body), nil
	}
	err = json.Unmarshal(body, responseType)
	if err != nil {
		logger.Default().Error("Error Unmarshaling JSON:", err)
		return nil, err
	}
	return responseType, nil

}
