/*
Package e2hformat is the formatter's package of the Enhanced Error Handling module
*/
package e2hformat

import (
	"fmt"
	"strings"

	"github.com/cdleo/go-commons/formatter"
	"github.com/cdleo/go-e2h"
)

type rawFormatter struct {
}

func newRawFormatter() Formatter {

	return &rawFormatter{}
}

func (s *rawFormatter) Source(err error) string {
	switch err := err.(type) {
	case e2h.EnhancedError:
		sourceError := err.Cause().Error()
		if len(err.Stack()) > 0 {
			stack := err.Stack()
			if len(stack[0].Message) > 0 {
				sourceError = fmt.Sprintf("%s [%s]", sourceError, stack[0].Message)
			}
		}
		return sourceError
	default:
		return err.Error()
	}
}

// This function returns the error stack information in a pretty format
func (s *rawFormatter) Format(err error, params Params) string {

	var result string
	var causeFormat, withInfoTrace, withoutInfoTrace string
	if params.Beautify {
		causeFormat = "%s\n"
		withInfoTrace = "%s (%s:%d)\n\t%s\n"
		withoutInfoTrace = "%s (%s:%d)\n"
	} else {
		causeFormat = "%s; "
		withInfoTrace = "%s (%s:%d) [%s]; "
		withoutInfoTrace = "%s (%s:%d); "
	}

	switch err := err.(type) {
	case e2h.EnhancedError:
		stackDetails := err.Stack()
		if params.InvertCallstack {
			for i := len(stackDetails) - 1; i >= 0; i-- {
				result += s.formatItem(withInfoTrace, withoutInfoTrace, params, stackDetails[i])
			}
			result += fmt.Sprintf(causeFormat, err.Cause().Error())
		} else {
			result = fmt.Sprintf(causeFormat, err.Cause().Error())
			for i := 0; i <= len(stackDetails)-1; i++ {
				stackItem := stackDetails[i]
				result += s.formatItem(withInfoTrace, withoutInfoTrace, params, stackItem)
			}
		}
	default:
		result = err.Error()
	}

	return strings.TrimSpace(result)
}

func (s *rawFormatter) formatItem(withInfoTrace string, withoutInfoTrace string, params Params, item e2h.StackDetails) string {

	filePath := formatter.FormatSourceFile(item.File, params.PathHidingMethod, params.PathHidingValue)

	if len(item.Message) > 0 {
		return fmt.Sprintf(withInfoTrace, item.FuncName, filePath, item.Line, item.Message)
	} else {
		return fmt.Sprintf(withoutInfoTrace, item.FuncName, filePath, item.Line)
	}
}
