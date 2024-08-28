package entity_const

import (
	"fmt"
)

type ValidationError struct {
	msg string
}

func NewValidationError(msg string) *ValidationError {
	return &ValidationError{msg: msg}
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error: %s", e.msg)
}

func (e *ValidationError) Unwrap() error {
	return fmt.Errorf("validation error: %s", e.msg)
}

func (e *ValidationError) Is(err error) bool {
	_, ok := err.(*ValidationError)
	return ok
}
