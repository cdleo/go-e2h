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
	Stack() []StackDetails
}

// Entity details with detailed info
type StackDetails struct {
	File     string
	Line     int
	FuncName string
	Message  string
}

// Entity enhancedError with error and details
type enhancedError struct {
	err   error
	stack []StackDetails
}

// This function returns the Error string plus the origin custom message (if exists)
func (e *enhancedError) Error() string {

	if len(e.stack[0].Message) > 0 {
		return fmt.Sprintf("%s: %s", e.err.Error(), e.stack[0].Message)
	}

	return e.err.Error()
}

// This function returns the source error
func (e *enhancedError) Cause() error {
	return e.err
}

// This function returns the callstack details
func (e *enhancedError) Stack() []StackDetails {
	return e.stack
}
