package errors

import "errors"

var (
	ErrInternalDatabaseError = errors.New("operation exec error")
	ErrAccountNotFound       = errors.New("account not found")
	ErrAccountDuplicated     = errors.New("document number duplicated")
)
