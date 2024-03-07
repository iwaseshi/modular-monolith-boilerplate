package error

import (
	"encoding/json"
	"fmt"
)

type ApiError struct {
	err     error
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewBusinessError(err error, messages ...fmt.Stringer) *ApiError {
	var msg string
	if len(messages) > 0 && messages[0] != nil {
		msg = messages[0].String()
	} else {
		msg = "Business error occurred"
	}

	return &ApiError{
		err:     err,
		Code:    400,
		Message: msg,
	}
}

func NewSystemError(err error, messages ...fmt.Stringer) *ApiError {
	var msg string
	if len(messages) > 0 && messages[0] != nil {
		msg = messages[0].String()
	} else {
		msg = "System error occurred"
	}

	return &ApiError{
		err:     err,
		Code:    500,
		Message: msg,
	}
}

func (e *ApiError) Error() string {
	errorJSON, err := json.Marshal(e)
	if err != nil {
		// エラーのマーシャリングに失敗した場合、フォールバックとして簡易的なエラーメッセージを返す
		return fmt.Sprintf("Error marshalling ApiError to JSON: %v", err)
	}
	return string(errorJSON)
}
