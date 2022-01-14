package e2h

import (
	"fmt"
	"runtime"
)

func TraceA(e error) error {
	return addTrace(e, "")
}
func Trace(e error, message string) error {
	return addTrace(e, message)
}
func Tracef(e error, format string, args ...interface{}) error {
	return addTrace(e, format, args...)
}

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
