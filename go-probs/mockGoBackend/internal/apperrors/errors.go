// Package apperrors defines domain-level errors for the mock backend.
//
// Go's error handling is explicit: functions return (T, error) instead of throwing.
// Callers decide whether to wrap, log, or propagate. Wrapping with fmt.Errorf("%w", err)
// preserves the original error for errors.Is / errors.As checks up the stack.
package apperrors

import "errors"

var (
	ErrNotFound     = errors.New("resource not found")
	ErrInvalidInput = errors.New("invalid input")
)
