package e2h_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cdleo/go-e2h"
)

func TestEnhancedError_StdError_GetCause(t *testing.T) {

	// Setup
	stdErr := fmt.Errorf("This is a standard error")

	output := e2h.Cause(stdErr)

	require.Equal(t, output, "This is a standard error")
}

func TestEnhancedError_FormatStdError_Pretty(t *testing.T) {

	// Setup
	stdErr := fmt.Errorf("This is a standard error")

	output := e2h.FormatPretty(stdErr)

	require.Equal(t, output, "This is a standard error")
}

func TestEnhancedError_FormatStdError_JSON(t *testing.T) {

	// Setup
	stdErr := fmt.Errorf("This is a standard error")

	outputJSON := e2h.FormatJSON(stdErr)

	require.Equal(t, string(outputJSON[:]), "{\"error\":\"This is a standard error\",\"stack_trace\":[]}")
}

func TestEnhancedError_EnhError_GetCause(t *testing.T) {

	// Setup
	enhancedErr := e2h.Trace(fmt.Errorf("This is a standard error"), "Error wrapped with additional info")

	output := e2h.Cause(enhancedErr)

	require.Equal(t, output, "This is a standard error")
}

func TestEnhancedError_FormatEnhError_Pretty(t *testing.T) {

	// Setup
	enhancedErr := e2h.Trace(fmt.Errorf("This is a standard error"), "Error wrapped with additional info")

	output := e2h.FormatPretty(enhancedErr)

	require.Equal(t, output, "- Error wrapped with additional info:\n  github.com/cdleo/go-e2h_test.TestEnhancedError_FormatEnhError_Pretty\n  \t/home/christian/sources/src/github.com/cdleo/go-e2h/e2h_test.go:55\n- This is a standard error\n")
}

func TestEnhancedError_FormatEnhError_JSON(t *testing.T) {

	// Setup
	enhancedErr := e2h.Trace(fmt.Errorf("This is a standard error"), "Error wrapped with additional info")

	outputJSON := e2h.FormatJSON(enhancedErr)

	require.Equal(t, string(outputJSON[:]), "{\"error\":\"This is a standard error: Error wrapped with additional info\",\"stack_trace\":[{\"func\":\"github.com/cdleo/go-e2h_test.TestEnhancedError_FormatEnhError_JSON\",\"caller\":\"/home/christian/sources/src/github.com/cdleo/go-e2h/e2h_test.go:65\",\"context\":\"Error wrapped with additional info\"}]}")
}

func TestEnhancedError_EnhError_StackTraces(t *testing.T) {

	// Setup
	enhancedErr := e2h.Trace(fmt.Errorf("This is a standard error"), "Error wrapped with additional info")
	enhancedErr = e2h.Tracef(enhancedErr, "This is the %dnd. stack level", 2)
	enhancedErr = e2h.TraceA(enhancedErr)

	output := e2h.FormatPretty(enhancedErr)

	require.Equal(t, output, "  github.com/cdleo/go-e2h_test.TestEnhancedError_EnhError_StackTraces\n  \t/home/christian/sources/src/github.com/cdleo/go-e2h/e2h_test.go:77\n- This is the 2nd. stack level:\n  github.com/cdleo/go-e2h_test.TestEnhancedError_EnhError_StackTraces\n  \t/home/christian/sources/src/github.com/cdleo/go-e2h/e2h_test.go:76\n- Error wrapped with additional info:\n  github.com/cdleo/go-e2h_test.TestEnhancedError_EnhError_StackTraces\n  \t/home/christian/sources/src/github.com/cdleo/go-e2h/e2h_test.go:75\n- This is a standard error\n")
}
