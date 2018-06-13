package goerr

import "fmt"

// StringerError extends Go's error interface to add a verbose message
type StringerError interface {
	// Error() reports the go-style error string
	error

	// String reports text describing the error in greater detail. Multiple-line
	// error messages are allowed. Return value should always end with a line
	// break.
	fmt.Stringer
}

type BasicStringerError struct {
	err error  // Go-style error string
	msg string // Verbose error string
}

func New(err error, message string) BasicStringerError {
	return BasicStringerError{err, message}
}
