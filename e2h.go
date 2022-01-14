package e2h

import (
	"fmt"
)

type EnhancedError interface {
	Error() string
	Cause() error
}

type details struct {
	file     string
	line     int
	funcName string
	message  string
}

type enhancedError struct {
	err   error
	stack []details
}

func (e *enhancedError) Error() string {

	if e.stack[0].message != "" {
		return fmt.Sprintf("%s: %s", e.err.Error(), e.stack[0].message)
	}

	return e.err.Error()
}

func (e *enhancedError) Cause() error {
	return e.err
}
