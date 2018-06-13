package errlist

import (
	"fmt"
	"strings"
)

type ErrList []error

// New creates an empty error list.
func New() *ErrList {
	list := ErrList([]error{})
	return &list
}

// Use returns the passed error list if it is not nil, otherwise it creates a
// new error list. This is a useful shorthand when creating a function that may
// take from the caller either an error list or nil.
func Use(passed *ErrList) *ErrList {
	if passed == nil {
		return New()
	}
	return passed
}

// Add appends the given error to the error list if and only if the error is not
// nil. If the given error was not nil, this method returns true.
func (el *ErrList) Add(err error) bool {
	if err != nil {
		*el = append(*el, err)
		return true
	}
	return false
}

// Get reports this error list, which is a valid error implementor, if and only
// if the length of the error list is greater than zero. Otherwise, nil is
// reported. Use this when reporting the error list to the caller.
func (el *ErrList) Get() error {
	if len(*el) == 0 {
		return nil
	}
	return el
}

// Error returns a string that aggregates error strings from each error in the
// list.
func (el *ErrList) Error() string {
	// Special case: error list length of 0 reports "no error"
	if len(*el) == 0 {
		return "no error"
	}

	// Special case: error list length of 1 adds no extra formatting
	if len(*el) == 1 {
		return (*el)[0].Error()
	}

	// Normal case: error list with multiple errors reporst a formatted string
	messages := []string{}
	for i, err := range *el {
		msg := err.Error()
		messages = append(messages, fmt.Sprintf("(error %d: %s)", i, msg))
	}

	return "error list: " + strings.Join(messages, ", ")
}

func (el *ErrList) String() string {
	// Special case: error list length of 0 reports "no error"
	if len(*el) == 0 {
		return "No errors in list."
	}

	// Special case: error list length of 1 adds no extra formatting
	if len(*el) == 1 {
		return "One error in list: " + (*el)[0].Error() + "\n"
	}

	output := "Multiple errors:\n"

	for i, err := range *el {
		msg := err.Error()
		output += "-" + msg + "\n"
	}

}
