package utils

import "errors"

// Define custom error messages for the application.
var (
	ErrInvalidCommand = errors.New("invalid command")
	ErrInvalidArgs    = errors.New("invalid arguments")
	ErrInvalidID      = errors.New("invalid task ID")
	ErrTaskNotFound   = errors.New("task not found")
	ErrInvalidStatus  = errors.New("invalid status")
)
