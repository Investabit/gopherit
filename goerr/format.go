package goerr

import "fmt"

type ErrorFormatter struct {
	format string
	extra  []interface{}
}

// FormatError formats a Go error interface
func NewErrorFormatter(format string, extra ...interface{}) ErrorFormatter {
	return ErrorFormatter{
		format, extra,
	}
}

func (formatter ErrorFormatter) Format(err error, extra ...interface{}) error {
	extraFields := append(formatter.extra, extra...)

	fields := append([]interface{}{}, err.Error())
	fields = append(fields, extraFields...)
	return fmt.Errorf(formatter.format, fields)
}
