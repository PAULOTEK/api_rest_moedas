package errors

import (
	"errors"
	"fmt"
)

type NotFoundError struct {
	message string
}

func NewNotFoundError(message string) *NotFoundError {
	return &NotFoundError{
		message: message,
	}
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("Error: %s", e.message)
}

func IsNotFoundError(err error) bool {
	var notFoundError *NotFoundError
	ok := errors.As(err, &notFoundError)
	return ok
}
