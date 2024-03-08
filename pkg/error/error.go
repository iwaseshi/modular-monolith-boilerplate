package error

import (
	"encoding/json"
	"fmt"
	"strings"
)

type ApiError struct {
	err     error
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type defaultErrorMessage string

func (msg defaultErrorMessage) String() string {
	return string(msg)
}

func newApiError(err error, code int, messages ...fmt.Stringer) *ApiError {
	var msgParts []string
	for _, message := range messages {
		if message != nil {
			msgParts = append(msgParts, message.String())
		}
	}
	msg := ""
	if len(msgParts) > 0 {
		msg = strings.Join(msgParts, " ")
	}

	return &ApiError{
		err:     err,
		Code:    code,
		Message: msg,
	}
}

func NewBusinessError(err error, messages ...fmt.Stringer) *ApiError {
	if len(messages) == 0 {
		messages = append(messages, defaultErrorMessage("Business error occurred"))
	}
	return newApiError(err, 400, messages...)
}

func NewSystemError(err error, messages ...fmt.Stringer) *ApiError {
	if len(messages) == 0 {
		messages = append(messages, defaultErrorMessage("System error occurred"))
	}
	return newApiError(err, 500, messages...)
}

func (e *ApiError) Error() string {
	errorJSON, err := json.Marshal(e)
	if err != nil {
		// エラーのマーシャリングに失敗した場合、フォールバックとして簡易的なエラーメッセージを返す
		return fmt.Sprintf("Error marshalling ApiError to JSON: %v", err)
	}
	return string(errorJSON)
}

func (e *ApiError) Cause() error {
	return e.err
}
