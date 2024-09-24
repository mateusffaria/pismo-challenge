package errors

import (
	"fmt"
)

type bodyError struct {
	message interface{}
}

// TODO: Change to new format
func NewBodyError(msg interface{}) *bodyError {
	return &bodyError{
		message: msg,
	}
}

func (e *bodyError) Error() string {
	return fmt.Sprintf("invalid body request: %v", e.message)
}
