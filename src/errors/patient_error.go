package errors

import "fmt"

type PatientError struct {
	Code    int
	Message string
}

func (e *PatientError) Error() string {
	return fmt.Sprintf("Patient Error %d: %s", e.Code, e.Message)
}

func NewPatientError(code int, msg string) error {
	return &PatientError{
		Code:    code,
		Message: msg,
	}
}
