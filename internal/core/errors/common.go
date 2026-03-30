package core_errors

import "errors"

var (
	ErrNotFound        = errors.New("Not Found")
	ErrInvalidArgument = errors.New("Invalid Argument")
	ErrConflict        = errors.New("Conflict")
)
