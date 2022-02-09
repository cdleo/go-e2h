/*
Package e2h its the package of the Enhanced Error Handling module
*/
package e2h

import (
	"fmt"
)

// ExtendedError interface
type EnhancedError interface {
	Error() string
	Cause() error
}

// Entity details with detailed info
type details struct {
	file     string
	line     int
	funcName string
	message  string
}

// Entity enhancedError with error and details
type enhancedError struct {
	err   error
	stack []details
}

// This function returns the Error string plus the origin custom message (if exists)
func (e *enhancedError) Error() string {

	if len(e.stack[0].message) > 0 {
		return fmt.Sprintf("%s: %s", e.err.Error(), e.stack[0].message)
	}

	return e.err.Error()
}

// This function returns the source error
func (e *enhancedError) Cause() error {
	return e.err
}
