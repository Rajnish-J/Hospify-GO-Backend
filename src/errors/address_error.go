package errors

import "fmt"

type AddressError struct {
	Code    int
	Message string
}

func (e *AddressError) Error() string {
	return fmt.Sprintf("Address Error %d: %s", e.Code, e.Message)
}

func NewAddressError(code int, msg string) error {
	return &AddressError{
		Code:    code,
		Message: msg,
	}
}
