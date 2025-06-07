package errors

import "fmt"

type AppError struct {
	Code    int
	Message string
}

func (e *AppError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

func NewAppError(code int, msg string) error {
	return &AppError{
		Code:    code,
		Message: msg,
	}
}
