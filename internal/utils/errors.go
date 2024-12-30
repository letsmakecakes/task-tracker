package utils

import "errors"

var (
	ErrInvalidCommand = errors.New("invalid command")
	ErrInvalidArgs    = errors.New("invalid arguments")
	ErrInvalidID      = errors.New("invalid task ID")
	ErrInvalidStatus  = errors.New("invalid status")
)
