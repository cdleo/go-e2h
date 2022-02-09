/*
Package e2h its the package of the Enhanced Error Handling module
*/
package e2h

import (
	"encoding/json"
	"fmt"
	"strings"
)

type jsonStack struct {
	FuncName string `json:"func"`
	Caller   string `json:"caller"`
	Context  string `json:"context,omitempty"`
}

func newJSONStack(item *details, rootFolder2Show string) jsonStack {
	return jsonStack{
		FuncName: item.funcName,
		Caller:   fmt.Sprintf("%s:%d", removePathBeforeFolder(item.file, rootFolder2Show), item.line),
		Context:  item.message,
	}
}

type jsonDetails struct {
	Err   string      `json:"error"`
	Stack []jsonStack `json:"stack_trace"`
}

// This function returns an string containing the description of the very first error in the stack
func Source(err error) string {
	switch err.(type) {
	case EnhancedError:
		return err.(*enhancedError).Cause().Error()
	default:
		return err.Error()
	}
}

// This function returns the error stack information in a JSON format
func FormatJSON(err error, rootFolder2Show string, indented bool) []byte {

	details := jsonDetails{
		Err:   err.Error(),
		Stack: make([]jsonStack, 0),
	}

	switch err.(type) {
	case EnhancedError:
		stkError := err.(*enhancedError)
		for i := len(stkError.stack) - 1; i >= 0; i-- {
			details.Stack = append(details.Stack, newJSONStack(&stkError.stack[i], rootFolder2Show))
		}
	default:
		//Do Nothing
	}

	var result []byte
	var marshalError error
	if indented {
		result, marshalError = json.MarshalIndent(details, "", "\t")
	} else {
		result, marshalError = json.Marshal(details)
	}
	if marshalError != nil {
		return []byte{0}
	}
	return result
}

// This function returns the error stack information in a pretty format
func FormatPretty(err error, rootFolder2Show string) string {

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
			result += fmt.Sprintf("  \t%s:%d\n", removePathBeforeFolder(stackItem.file, rootFolder2Show), stackItem.line)
		}
		result += fmt.Sprintf("- %s\n", stkError.Cause())
	default:
		result = err.Error()
	}
	return result
}

// Utility funtion that removes the first part of the file path til the folder indicated in `newRootFolder` argument
func removePathBeforeFolder(file string, newRootFolder string) string {

	if len(newRootFolder) <= 0 {
		return file
	}

	fileParts := strings.Split(file, newRootFolder)
	if len(fileParts) < 2 {
		return file
	}
	return newRootFolder + fileParts[1]
}
