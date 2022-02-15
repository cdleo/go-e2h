/*
Package e2hformat is the formatter's package of the Enhanced Error Handling module
*/
package e2hformat

import (
	"encoding/json"
	"fmt"

	"github.com/cdleo/go-e2h"
)

type jsonStack struct {
	FuncName string `json:"func"`
	Caller   string `json:"caller"`
	Context  string `json:"context,omitempty"`
}

type jsonDetails struct {
	Err   string      `json:"error"`
	Stack []jsonStack `json:"stack_trace"`
}

type jsonSource struct {
	Err     string `json:"error"`
	Context string `json:"context,omitempty"`
}

func newJSONStack(item *e2h.StackDetails, hidingMethod HidingMethod, pathOrFolder string) jsonStack {

	filePath := formatSourceFile(item.File, hidingMethod, pathOrFolder)

	return jsonStack{
		FuncName: item.FuncName,
		Caller:   fmt.Sprintf("%s:%d", filePath, item.Line),
		Context:  item.Message,
	}
}

type jsonFormatter struct {
}

func newJSONFormatter() Formatter {

	return &jsonFormatter{}
}

func (s *jsonFormatter) Source(err error) string {

	var source jsonSource
	switch err := err.(type) {
	case e2h.EnhancedError:
		source.Err = err.Cause().Error()
		if len(err.Stack()) > 0 {
			stack := err.Stack()
			if len(stack[0].Message) > 0 {
				source.Context = stack[0].Message
			}
		}
	default:
		source.Err = err.Error()
	}

	if result, marshalError := json.Marshal(source); marshalError != nil {
		return ""
	} else {
		return string(result)
	}

}

// This function returns the error stack information in a JSON format
func (s *jsonFormatter) Format(err error, params Params) string {

	details := jsonDetails{
		Err:   err.Error(),
		Stack: make([]jsonStack, 0),
	}

	switch err := err.(type) {
	case e2h.EnhancedError:
		details.Err = err.Cause().Error()
		stackDetails := err.Stack()
		if params.InvertCallstack {
			for i := len(stackDetails) - 1; i >= 0; i-- {
				details.Stack = append(details.Stack, newJSONStack(&stackDetails[i],
					params.PathHidingMethod, params.PathHidingValue))
			}
		} else {
			for i := 0; i <= len(stackDetails)-1; i++ {
				details.Stack = append(details.Stack, newJSONStack(&stackDetails[i],
					params.PathHidingMethod, params.PathHidingValue))
			}
		}

	default:
		//Do Nothing
	}

	var result []byte
	var marshalError error
	if params.Beautify {
		result, marshalError = json.MarshalIndent(details, "", "\t")
	} else {
		result, marshalError = json.Marshal(details)
	}
	if marshalError != nil {
		return ""
	}
	return string(result)
}
