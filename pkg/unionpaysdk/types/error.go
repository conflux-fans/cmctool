package types

import "fmt"

type ApiError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (e *ApiError) Error() string {
	return fmt.Sprintf("code: %s, message: %s, data: %v", e.Code, e.Message, e.Data)
}
