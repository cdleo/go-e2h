/*
Package e2h_test its the test package of the Enhanced Error Handling module
*/
package e2h_test

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/cdleo/go-e2h"
	e2hformat "github.com/cdleo/go-e2h/formatter"
	"github.com/stretchr/testify/require"
)

func TestEnhancedError_StdErr_RawFormatter_GetSource(t *testing.T) {

	// Setup
	stdErr := fmt.Errorf("This is a standard error")
	rawFormatter, err := e2hformat.NewFormatter(e2hformat.Format_Raw)
	require.Nil(t, err)

	// Execute
	output_raw := rawFormatter.Source(stdErr)

	// Check
	require.Equal(t, output_raw, stdErr.Error())
}

func TestEnhancedError_StdErr_JSONFormatter_GetSource(t *testing.T) {

	// Setup
	stdErr := fmt.Errorf("This is a standard error")
	jsonFormatter, err := e2hformat.NewFormatter(e2hformat.Format_JSON)
	require.Nil(t, err)

	// Execute
	output_json := jsonFormatter.Source(stdErr)

	// Check
	require.Equal(t, output_json, "{\"error\":\"This is a standard error\"}")
}

func TestEnhancedError_StdErr_RawFormatter_Format(t *testing.T) {

	// Setup
	stdErr := fmt.Errorf("This is a standard error")
	rawFormatter, _ := e2hformat.NewFormatter(e2hformat.Format_Raw)
	params := e2hformat.Params{}

	// Execute
	output := rawFormatter.Format(stdErr, params)

	// Check
	require.Equal(t, output, "This is a standard error")
}

func TestEnhancedError_StdErr_JSONFormatter_Format(t *testing.T) {

	// Setup
	stdErr := fmt.Errorf("This is a standard error")
	jsonFormatter, _ := e2hformat.NewFormatter(e2hformat.Format_JSON)
	params := e2hformat.Params{}

	// Execute
	outputJSON := jsonFormatter.Format(stdErr, params)

	// Check
	require.Equal(t, outputJSON, "{\"error\":\"This is a standard error\",\"stack_trace\":[]}")
}

func TestEnhancedError_EnhErr_RawFormatter_GetCause(t *testing.T) {

	// Setup
	enhancedErr := e2h.Tracem(fmt.Errorf("This is a standard error"), "Error wrapped with additional info")
	rawFormatter, _ := e2hformat.NewFormatter(e2hformat.Format_Raw)

	// Execute
	output := rawFormatter.Source(enhancedErr)

	// Check
	require.Equal(t, output, "This is a standard error [Error wrapped with additional info]")
}

func TestEnhancedError_EnhErr_JSONFormatter_GetCause(t *testing.T) {

	// Setup
	enhancedErr := e2h.Tracem(fmt.Errorf("This is a standard error"), "Error wrapped with additional info")
	jsonFormatter, _ := e2hformat.NewFormatter(e2hformat.Format_JSON)

	// Execute
	output := jsonFormatter.Source(enhancedErr)

	// Check
	require.Equal(t, output, "{\"error\":\"This is a standard error\",\"context\":\"Error wrapped with additional info\"}")
}

func TestEnhancedError_EnhErr_RawFormatter_Format(t *testing.T) {

	// Setup
	enhancedErr := e2h.Tracem(fmt.Errorf("This is a standard error"), "Error wrapped with additional info")
	rawFormatter, _ := e2hformat.NewFormatter(e2hformat.Format_Raw)
	params := e2hformat.Params{}

	// Execute
	output := rawFormatter.Format(enhancedErr, params)

	// Check
	require.Equal(t, output, "This is a standard error; github.com/cdleo/go-e2h_test.TestEnhancedError_EnhErr_RawFormatter_Format (/home/christian/sources/src/github.com/cdleo/go-e2h/e2h_test.go:104) [Error wrapped with additional info];")
}

func TestEnhancedError_EnhErr_JSONFormatter_Format(t *testing.T) {

	// Setup
	enhancedErr := e2h.Tracem(fmt.Errorf("This is a standard error"), "Error wrapped with additional info")
	jsonFormatter, _ := e2hformat.NewFormatter(e2hformat.Format_JSON)
	params := e2hformat.Params{}

	// Execute
	output := jsonFormatter.Format(enhancedErr, params)

	// Check
	require.Equal(t, output, "{\"error\":\"This is a standard error\",\"stack_trace\":[{\"func\":\"github.com/cdleo/go-e2h_test.TestEnhancedError_EnhErr_JSONFormatter_Format\",\"caller\":\"/home/christian/sources/src/github.com/cdleo/go-e2h/e2h_test.go:118\",\"context\":\"Error wrapped with additional info\"}]}")
}

func TestEnhancedError_EnhErr_RawFormatter_Format_Beautified(t *testing.T) {

	// Setup
	enhancedErr := e2h.Tracem(fmt.Errorf("This is a standard error"), "Error wrapped with additional info")
	rawFormatter, _ := e2hformat.NewFormatter(e2hformat.Format_Raw)
	params := e2hformat.Params{
		Beautify: true,
	}

	// Execute
	output := rawFormatter.Format(enhancedErr, params)

	// Check
	require.Equal(t, output, "This is a standard error\ngithub.com/cdleo/go-e2h_test.TestEnhancedError_EnhErr_RawFormatter_Format_Beautified (/home/christian/sources/src/github.com/cdleo/go-e2h/e2h_test.go:132)\n\tError wrapped with additional info")
}

func TestEnhancedError_EnhErr_JSONFormatter_Format_Beautified(t *testing.T) {

	// Setup
	enhancedErr := e2h.Tracem(fmt.Errorf("This is a standard error"), "Error wrapped with additional info")
	jsonFormatter, _ := e2hformat.NewFormatter(e2hformat.Format_JSON)
	params := e2hformat.Params{
		Beautify: true,
	}

	// Execute
	output := jsonFormatter.Format(enhancedErr, params)

	// Check
	require.Equal(t, output, "{\n\t\"error\": \"This is a standard error\",\n\t\"stack_trace\": [\n\t\t{\n\t\t\t\"func\": \"github.com/cdleo/go-e2h_test.TestEnhancedError_EnhErr_JSONFormatter_Format_Beautified\",\n\t\t\t\"caller\": \"/home/christian/sources/src/github.com/cdleo/go-e2h/e2h_test.go:148\",\n\t\t\t\"context\": \"Error wrapped with additional info\"\n\t\t}\n\t]\n}")
}

func TestEnhancedError_EnhErr_RawFormatter_Format_Inverted(t *testing.T) {

	// Setup
	enhancedErr := e2h.Tracem(fmt.Errorf("This is a standard error"), "Error wrapped with additional info")
	rawFormatter, _ := e2hformat.NewFormatter(e2hformat.Format_Raw)
	params := e2hformat.Params{
		InvertCallstack: true,
	}

	// Execute
	output := rawFormatter.Format(enhancedErr, params)

	// Check
	require.Equal(t, output, "github.com/cdleo/go-e2h_test.TestEnhancedError_EnhErr_RawFormatter_Format_Inverted (/home/christian/sources/src/github.com/cdleo/go-e2h/e2h_test.go:164) [Error wrapped with additional info]; This is a standard error;")
}

func TestEnhancedError_EnhErr_JSONFormatter_Format_Inverted(t *testing.T) {

	// Setup
	enhancedErr := e2h.Tracem(fmt.Errorf("This is a standard error"), "Error wrapped with additional info")
	jsonFormatter, _ := e2hformat.NewFormatter(e2hformat.Format_JSON)
	params := e2hformat.Params{
		InvertCallstack: true,
	}

	// Execute
	output := jsonFormatter.Format(enhancedErr, params)

	// Check
	require.Equal(t, output, "{\"error\":\"This is a standard error\",\"stack_trace\":[{\"func\":\"github.com/cdleo/go-e2h_test.TestEnhancedError_EnhErr_JSONFormatter_Format_Inverted\",\"caller\":\"/home/christian/sources/src/github.com/cdleo/go-e2h/e2h_test.go:180\",\"context\":\"Error wrapped with additional info\"}]}")
}

func TestEnhancedError_EnhErr_RawFormatter_Format_FullPathHidden(t *testing.T) {

	// Setup
	enhancedErr := e2h.Tracem(fmt.Errorf("This is a standard error"), "Error wrapped with additional info")
	rawFormatter, _ := e2hformat.NewFormatter(e2hformat.Format_Raw)
	_, b, _, _ := runtime.Caller(0)
	hideThisPath := filepath.Dir(b) + string(os.PathSeparator)
	params := e2hformat.Params{
		PathHidingMethod: e2hformat.HidingMethod_FullBaseline,
		PathHidingValue:  hideThisPath,
	}

	// Execute
	output := rawFormatter.Format(enhancedErr, params)

	// Check
	require.Equal(t, output, "This is a standard error; github.com/cdleo/go-e2h_test.TestEnhancedError_EnhErr_RawFormatter_Format_FullPathHidden (e2h_test.go:196) [Error wrapped with additional info];")
}

func TestEnhancedError_EnhErr_JSONFormatter_Format_FullPathHidden(t *testing.T) {

	// Setup
	enhancedErr := e2h.Tracem(fmt.Errorf("This is a standard error"), "Error wrapped with additional info")
	jsonFormatter, _ := e2hformat.NewFormatter(e2hformat.Format_JSON)
	_, b, _, _ := runtime.Caller(0)
	hideThisPath := filepath.Dir(b) + string(os.PathSeparator)
	params := e2hformat.Params{
		PathHidingMethod: e2hformat.HidingMethod_FullBaseline,
		PathHidingValue:  hideThisPath,
	}

	// Execute
	output := jsonFormatter.Format(enhancedErr, params)

	// Check
	require.Equal(t, output, "{\"error\":\"This is a standard error\",\"stack_trace\":[{\"func\":\"github.com/cdleo/go-e2h_test.TestEnhancedError_EnhErr_JSONFormatter_Format_FullPathHidden\",\"caller\":\"e2h_test.go:215\",\"context\":\"Error wrapped with additional info\"}]}")
}

func TestEnhancedError_EnhErr_RawFormatter_Format_PartialPathHidden(t *testing.T) {

	// Setup
	enhancedErr := e2h.Tracem(fmt.Errorf("This is a standard error"), "Error wrapped with additional info")
	rawFormatter, _ := e2hformat.NewFormatter(e2hformat.Format_Raw)
	_, b, _, _ := runtime.Caller(0)
	path := strings.Split(filepath.Dir(b), string(os.PathSeparator))
	params := e2hformat.Params{
		PathHidingMethod: e2hformat.HidingMethod_ToFolder,
		PathHidingValue:  path[len(path)-1],
	}

	// Execute
	output := rawFormatter.Format(enhancedErr, params)

	// Check
	require.Equal(t, output, "This is a standard error; github.com/cdleo/go-e2h_test.TestEnhancedError_EnhErr_RawFormatter_Format_PartialPathHidden (go-e2h/e2h_test.go:234) [Error wrapped with additional info];")
}

func TestEnhancedError_EnhErr_JSONFormatter_Format_PartialPathHidden(t *testing.T) {

	// Setup
	enhancedErr := e2h.Tracem(fmt.Errorf("This is a standard error"), "Error wrapped with additional info")
	jsonFormatter, _ := e2hformat.NewFormatter(e2hformat.Format_JSON)
	_, b, _, _ := runtime.Caller(0)
	path := strings.Split(filepath.Dir(b), string(os.PathSeparator))
	params := e2hformat.Params{
		PathHidingMethod: e2hformat.HidingMethod_ToFolder,
		PathHidingValue:  path[len(path)-1],
	}

	// Execute
	output := jsonFormatter.Format(enhancedErr, params)

	// Check
	require.Equal(t, output, "{\"error\":\"This is a standard error\",\"stack_trace\":[{\"func\":\"github.com/cdleo/go-e2h_test.TestEnhancedError_EnhErr_JSONFormatter_Format_PartialPathHidden\",\"caller\":\"go-e2h/e2h_test.go:253\",\"context\":\"Error wrapped with additional info\"}]}")
}

func TestEnhancedError_EnhError_JSONFormatter_Format_MultipleTraces(t *testing.T) {

	// Setup
	enhancedErr := e2h.Tracem(fmt.Errorf("This is a standard error"), "Error wrapped with additional info")
	enhancedErr = e2h.Tracef(enhancedErr, "This is the %dnd. stack level", 2)
	enhancedErr = e2h.Trace(enhancedErr)
	jsonFormatter, _ := e2hformat.NewFormatter(e2hformat.Format_JSON)
	_, b, _, _ := runtime.Caller(0)
	hideThisPath := filepath.Dir(b) + string(os.PathSeparator)
	params := e2hformat.Params{
		Beautify:         true,
		PathHidingMethod: e2hformat.HidingMethod_FullBaseline,
		PathHidingValue:  hideThisPath,
	}

	// Execute
	output := jsonFormatter.Format(enhancedErr, params)

	// Check
	require.Equal(t, output, "{\n\t\"error\": \"This is a standard error\",\n\t\"stack_trace\": [\n\t\t{\n\t\t\t\"func\": \"github.com/cdleo/go-e2h_test.TestEnhancedError_EnhError_JSONFormatter_Format_MultipleTraces\",\n\t\t\t\"caller\": \"e2h_test.go:272\",\n\t\t\t\"context\": \"Error wrapped with additional info\"\n\t\t},\n\t\t{\n\t\t\t\"func\": \"github.com/cdleo/go-e2h_test.TestEnhancedError_EnhError_JSONFormatter_Format_MultipleTraces\",\n\t\t\t\"caller\": \"e2h_test.go:273\",\n\t\t\t\"context\": \"This is the 2nd. stack level\"\n\t\t},\n\t\t{\n\t\t\t\"func\": \"github.com/cdleo/go-e2h_test.TestEnhancedError_EnhError_JSONFormatter_Format_MultipleTraces\",\n\t\t\t\"caller\": \"e2h_test.go:274\"\n\t\t}\n\t]\n}")
}
