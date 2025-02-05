package errors

import "fmt"

type ErrorKind string

const (
	NotFound ErrorKind = "not_found"
	Invalid  ErrorKind = "invalid"
	Internal ErrorKind = "internal"
)

type CustomError struct {
	Kind    ErrorKind
	Message string
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("[%s] %s", e.Kind, e.Message)
}

func NewCustomError(kind ErrorKind, message string) error {
	return &CustomError{
		Kind:    kind,
		Message: message,
	}
}
