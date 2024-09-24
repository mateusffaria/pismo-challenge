package errors

import "errors"

var (
	ErrInternalDatabaseError = errors.New("operation exec error")
	ErrNotFound              = errors.New("operation type not found")
)
