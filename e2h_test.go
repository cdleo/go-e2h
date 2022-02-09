/*
Package e2h_test its the test package of the Enhanced Error Handling module
*/
package e2h_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cdleo/go-e2h"
)

func TestEnhancedError_StdError_GetSource(t *testing.T) {

	// Setup
	stdErr := fmt.Errorf("This is a standard error")

	// Execute
	output := e2h.Source(stdErr)

	// Check
	require.Equal(t, output, "This is a standard error")
}

func TestEnhancedError_FormatStdError_Pretty(t *testing.T) {

	// Setup
	stdErr := fmt.Errorf("This is a standard error")

	// Execute
	output := e2h.FormatPretty(stdErr, "")

	// Check
	require.Equal(t, output, "This is a standard error")
}

func TestEnhancedError_FormatStdError_JSON(t *testing.T) {

	// Setup
	stdErr := fmt.Errorf("This is a standard error")

	// Execute
	outputJSON := e2h.FormatJSON(stdErr, "", false)

	// Check
	require.Equal(t, string(outputJSON[:]), "{\"error\":\"This is a standard error\",\"stack_trace\":[]}")
}

func TestEnhancedError_EnhError_GetCause(t *testing.T) {

	// Setup
	enhancedErr := e2h.Tracem(fmt.Errorf("This is a standard error"), "Error wrapped with additional info")

	// Execute
	output := e2h.Source(enhancedErr)

	// Check
	require.Equal(t, output, "This is a standard error")
}

func TestEnhancedError_FormatEnhError_Pretty(t *testing.T) {

	// Setup
	enhancedErr := e2h.Tracem(fmt.Errorf("This is a standard error"), "Error wrapped with additional info")

	// Execute
	output := e2h.FormatPretty(enhancedErr, "github.com")

	// Check
	require.Equal(t, output, "- Error wrapped with additional info:\n  github.com/cdleo/go-e2h_test.TestEnhancedError_FormatEnhError_Pretty\n  \tgithub.com/cdleo/go-e2h/e2h_test.go:66\n- This is a standard error\n")
}

func TestEnhancedError_FormatEnhError_JSON_STD(t *testing.T) {

	// Setup
	enhancedErr := e2h.Tracem(fmt.Errorf("This is a standard error"), "Error wrapped with additional info")

	// Execute
	outputJSON := e2h.FormatJSON(enhancedErr, "github.com", false)

	// Check
	require.Equal(t, string(outputJSON[:]), "{\"error\":\"This is a standard error: Error wrapped with additional info\",\"stack_trace\":[{\"func\":\"github.com/cdleo/go-e2h_test.TestEnhancedError_FormatEnhError_JSON_STD\",\"caller\":\"github.com/cdleo/go-e2h/e2h_test.go:78\",\"context\":\"Error wrapped with additional info\"}]}")
}

func TestEnhancedError_FormatEnhError_JSON_Pretty(t *testing.T) {

	// Setup
	enhancedErr := e2h.Tracem(fmt.Errorf("This is a standard error"), "Error wrapped with additional info")

	// Execute
	outputJSON := e2h.FormatJSON(enhancedErr, "github.com", true)

	// Check
	require.Equal(t, string(outputJSON[:]), "{\n\t\"error\": \"This is a standard error: Error wrapped with additional info\",\n\t\"stack_trace\": [\n\t\t{\n\t\t\t\"func\": \"github.com/cdleo/go-e2h_test.TestEnhancedError_FormatEnhError_JSON_Pretty\",\n\t\t\t\"caller\": \"github.com/cdleo/go-e2h/e2h_test.go:90\",\n\t\t\t\"context\": \"Error wrapped with additional info\"\n\t\t}\n\t]\n}")
}

func TestEnhancedError_EnhError_StackTraces(t *testing.T) {

	// Setup
	enhancedErr := e2h.Tracem(fmt.Errorf("This is a standard error"), "Error wrapped with additional info")
	enhancedErr = e2h.Tracef(enhancedErr, "This is the %dnd. stack level", 2)
	enhancedErr = e2h.Trace(enhancedErr)

	// Execute
	output := e2h.FormatPretty(enhancedErr, "github.com")

	// Check
	require.Equal(t, output, "  github.com/cdleo/go-e2h_test.TestEnhancedError_EnhError_StackTraces\n  \tgithub.com/cdleo/go-e2h/e2h_test.go:104\n- This is the 2nd. stack level:\n  github.com/cdleo/go-e2h_test.TestEnhancedError_EnhError_StackTraces\n  \tgithub.com/cdleo/go-e2h/e2h_test.go:103\n- Error wrapped with additional info:\n  github.com/cdleo/go-e2h_test.TestEnhancedError_EnhError_StackTraces\n  \tgithub.com/cdleo/go-e2h/e2h_test.go:102\n- This is a standard error\n")
}
