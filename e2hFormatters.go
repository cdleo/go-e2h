package e2h

import (
	"encoding/json"
	"fmt"
)

type jsonStack struct {
	FuncName string `json:"func"`
	Caller   string `json:"caller"`
	Context  string `json:"context,omitempty"`
}

func newJSONStack(item *details) jsonStack {
	return jsonStack{
		FuncName: item.funcName,
		Caller:   fmt.Sprintf("%s:%d", item.file, item.line),
		Context:  item.message,
	}
}

type jsonDetails struct {
	Err   string      `json:"error"`
	Stack []jsonStack `json:"stack_trace"`
}

func Cause(err error) string {
	switch err.(type) {
	case EnhancedError:
		return err.(*enhancedError).Cause().Error()
	default:
		return err.Error()
	}
}

func FormatJSON(err error) []byte {

	details := jsonDetails{
		Err:   err.Error(),
		Stack: make([]jsonStack, 0),
	}

	switch err.(type) {
	case EnhancedError:
		stkError := err.(*enhancedError)
		for i := len(stkError.stack) - 1; i >= 0; i-- {
			details.Stack = append(details.Stack, newJSONStack(&stkError.stack[i]))
		}
	default:
		//Do Nothing
	}

	result, marshalError := json.Marshal(details)
	if marshalError != nil {
		return []byte{0}
	}
	return result
}

func FormatPretty(err error) string {

	var result string
	switch err.(type) {
	case EnhancedError:
		stkError := err.(*enhancedError)
		for i := len(stkError.stack) - 1; i >= 0; i-- {
			stackItem := stkError.stack[i]
			if len(stackItem.message) > 0 {
				result += fmt.Sprintf("- %s:\n", stackItem.message)
			}
			result += fmt.Sprintf("  %s\n", stackItem.funcName)
			result += fmt.Sprintf("  \t%s:%d\n", stackItem.file, stackItem.line)
		}
		result += fmt.Sprintf("- %s\n", stkError.Cause())
	default:
		result = err.Error()
	}
	return result
}
