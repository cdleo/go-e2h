/*
Package e2h its the package of the Enhanced Error Handling module
*/
package e2h

import (
	"fmt"
	"runtime"
)

// This function calls the addTrace in order to create or add stack info
func Trace(e error) error {
	return addTrace(e, "")
}

// Same as Trace, but adding a descriptive message
func Tracem(e error, message string) error {
	return addTrace(e, message)
}

// Same as Tracem, but the descriptive message can have formatted values
func Tracef(e error, format string, args ...interface{}) error {
	return addTrace(e, format, args...)
}

// This is the private function that creates the first EnhancedError
// with info or add the new info to the existing one
func addTrace(err error, format string, args ...interface{}) error {

	if err == nil {
		return nil
	}

	message := format
	if args != nil {
		message = fmt.Sprintf(format, args...)
	}
	pc, file, line, _ := runtime.Caller(2)
	info := details{
		file:     file,
		line:     line,
		funcName: runtime.FuncForPC(pc).Name(),
		message:  message,
	}

	switch err.(type) {
	case EnhancedError:
		err.(*enhancedError).stack = append(err.(*enhancedError).stack, info)
		return err

	default:
		return &enhancedError{
			err:   err,
			stack: append(make([]details, 0), info),
		}
	}
}
