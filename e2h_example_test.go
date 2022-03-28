/*
Package e2h_test its the test package of the Enhanced Error Handling module
*/
package e2h_test

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/cdleo/go-commons/formatter"
	"github.com/cdleo/go-e2h"
	e2hformat "github.com/cdleo/go-e2h/formatter"
)

func foo() error {

	//Doing something, an error it's returned
	err := fmt.Errorf("TheError")
	if err != nil {
		//It's not mandatory, but recommended to call TraceX() from deepest possible in the stack
		//to get the most additional data
		return e2h.Trace(err)
	}

	//Do stuff

	return nil
}

func bar() error {
	return e2h.Tracem(foo(), "Error executing foo()")
}

func ExampleEnhancedError() {
	_, b, _, _ := runtime.Caller(0)
	hideThisPath := filepath.Dir(b) + string(os.PathSeparator)

	params := e2hformat.Params{
		Beautify:         false,
		InvertCallstack:  true,
		PathHidingMethod: formatter.HidingMethod_FullBaseline,
		PathHidingValue:  hideThisPath,
	}

	err := e2h.Tracef(bar(), "Error executing [%s] function", "bar()")

	fmt.Printf("As Error => %v\n\n", err)

	fmt.Printf("**** Raw Formatter ****\n\n")

	rawFormatter, _ := e2hformat.NewFormatter(e2hformat.Format_Raw)

	fmt.Printf("Just cause => %s\n\n", rawFormatter.Source(err))

	fmt.Printf("Full info (inverted stack) =>\n%s\n\n", rawFormatter.Format(err, params))

	params.Beautify = true
	fmt.Printf("Full info (beautified / inverted stack) =>\n%s\n\n", rawFormatter.Format(err, params))

	fmt.Printf("**** JSON Formatter ****\n\n")

	jsonFormatter, _ := e2hformat.NewFormatter(e2hformat.Format_JSON)

	fmt.Printf("Just cause => %s\n\n", jsonFormatter.Source(err))

	params.InvertCallstack = false
	params.Beautify = false
	fmt.Printf("Full info =>\n%s\n\n", jsonFormatter.Format(err, params))

	params.Beautify = true
	fmt.Printf("Full info (beautified) =>\n%s\n", jsonFormatter.Format(err, params))

	// Output:
	// As Error => TheError
	//
	// **** Raw Formatter ****
	//
	// Just cause => TheError
	//
	// Full info (inverted stack) =>
	// github.com/cdleo/go-e2h_test.ExampleEnhancedError (e2h_example_test.go:47) [Error executing [bar()] function]; github.com/cdleo/go-e2h_test.bar (e2h_example_test.go:33) [Error executing foo()]; github.com/cdleo/go-e2h_test.foo (e2h_example_test.go:24); TheError;
	//
	// Full info (beautified / inverted stack) =>
	// github.com/cdleo/go-e2h_test.ExampleEnhancedError (e2h_example_test.go:47)
	// 	Error executing [bar()] function
	// github.com/cdleo/go-e2h_test.bar (e2h_example_test.go:33)
	// 	Error executing foo()
	// github.com/cdleo/go-e2h_test.foo (e2h_example_test.go:24)
	// TheError
	//
	// **** JSON Formatter ****
	//
	// Just cause => {"error":"TheError"}
	//
	// Full info =>
	// {"error":"TheError","stack_trace":[{"func":"github.com/cdleo/go-e2h_test.foo","caller":"e2h_example_test.go:24"},{"func":"github.com/cdleo/go-e2h_test.bar","caller":"e2h_example_test.go:33","context":"Error executing foo()"},{"func":"github.com/cdleo/go-e2h_test.ExampleEnhancedError","caller":"e2h_example_test.go:47","context":"Error executing [bar()] function"}]}
	//
	// Full info (beautified) =>
	// {
	// 	"error": "TheError",
	// 	"stack_trace": [
	// 		{
	// 			"func": "github.com/cdleo/go-e2h_test.foo",
	// 			"caller": "e2h_example_test.go:24"
	// 		},
	// 		{
	// 			"func": "github.com/cdleo/go-e2h_test.bar",
	// 			"caller": "e2h_example_test.go:33",
	// 			"context": "Error executing foo()"
	// 		},
	// 		{
	// 			"func": "github.com/cdleo/go-e2h_test.ExampleEnhancedError",
	// 			"caller": "e2h_example_test.go:47",
	// 			"context": "Error executing [bar()] function"
	// 		}
	// 	]
	// }

}
