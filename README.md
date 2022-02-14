# GO-E2H

[![Go Reference](https://pkg.go.dev/badge/github.com/cdleo/go-e2h/README.md.svg)](https://pkg.go.dev/github.com/cdleo/go-e2h/README.md)

[![Go Reference](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://pkg.go.dev/github.com/cdleo/go-e2h) [![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/cdleo/go-e2h/master/LICENSE) [![Build Status](https://travis-ci.org/rs/zerolog.svg?branch=master)](https://travis-ci.org/rs/zerolog) [![Coverage](http://gocover.io/_badge/github.com/rs/zerolog)](http://gocover.io/github.com/rs/zerolog)

GO Enhanced Error Handling (a.k.a. go-e2h) is a lightweight Golang module to add a better stack trace and context information on error events.

## General

We use an object of the provided interface EnhancedError to store the context and stack information:

```go
type EnhancedError interface {
    // This function returns the Error string plus the first context message (if exists)
	Error() string
    // This function returns the source error
	Cause() error
	// This function returns an array with the callstack details
	Stack() []StackDetails
}
```

**Note:** You will never need to use the `EnhancedError` specific type, due to all provided functions uses golang standard interface, which it's compatible.

To save the error callstack and add helpful information, we provide the following stateless functions:

```go
// This function just store or adds stack information
func Trace(e error) error

// Same as Trace, but adding a descriptive context message
func Tracem(e error, message string) error

// Same as Trace, but supporting a variadic function with format string as context information
func Tracef(e error, format string, args ...interface{}) error
```

Additionally, we provide a package called e2hformat in order to get the error information, over different formats:

```go
type Formatter interface {
	Source(err error) string
	Format(err error, params Params) string
}

type Params struct {
	Beautify         bool
	InvertCallstack  bool
	PathHidingMethod HidingMethod
	PathHidingValue  string
}


// This function returns an string containing the description of the very first error in the stack
func Source(err error) string 

// This function returns the error stack information in a JSON format
func FormatJSON(err error, rootFolder2Show string, indented bool) []byte

// This function returns the error stack information in a pretty format
func FormatPretty(err error, rootFolder2Show string) string 
```

## Usage

The use of this module is very simple, as you may see:
1 - You get/return a `standard GO error`.
2 - You call the `Trace` function (in all of yours versions) and get another error compatible object.
3 - At the highest point of the callstack, or the desired level to log/print the event, you call a formatter.

**Note:** You could call any of the `Trace` functions **even if no error have occurred**, in which case returns a `nil` value

The following example program shows the use of `Source`, `Trace` and `Formatting` options:
```go
package e2h_test

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

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
		PathHidingMethod: e2hformat.HidingMethod_FullBaseline,
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
	// github.com/cdleo/go-e2h_test.ExampleEnhancedError (e2h_example_test.go:46) [Error executing [bar()] function]; github.com/cdleo/go-e2h_test.bar (e2h_example_test.go:32) [Error executing foo()]; github.com/cdleo/go-e2h_test.foo (e2h_example_test.go:23); TheError;
	//
	// Full info (beautified / inverted stack) =>
	// github.com/cdleo/go-e2h_test.ExampleEnhancedError (e2h_example_test.go:46)
	// 	Error executing [bar()] function
	// github.com/cdleo/go-e2h_test.bar (e2h_example_test.go:32)
	// 	Error executing foo()
	// github.com/cdleo/go-e2h_test.foo (e2h_example_test.go:23)
	// TheError
	//
	// **** JSON Formatter ****
	//
	// Just cause => {"error":"TheError"}
	//
	// Full info =>
	// {"error":"TheError","stack_trace":[{"func":"github.com/cdleo/go-e2h_test.foo","caller":"e2h_example_test.go:23"},{"func":"github.com/cdleo/go-e2h_test.bar","caller":"e2h_example_test.go:32","context":"Error executing foo()"},{"func":"github.com/cdleo/go-e2h_test.ExampleEnhancedError","caller":"e2h_example_test.go:46","context":"Error executing [bar()] function"}]}
	//
	// Full info (beautified) =>
	// {
	// 	"error": "TheError",
	// 	"stack_trace": [
	// 		{
	// 			"func": "github.com/cdleo/go-e2h_test.foo",
	// 			"caller": "e2h_example_test.go:23"
	// 		},
	// 		{
	// 			"func": "github.com/cdleo/go-e2h_test.bar",
	// 			"caller": "e2h_example_test.go:32",
	// 			"context": "Error executing foo()"
	// 		},
	// 		{
	// 			"func": "github.com/cdleo/go-e2h_test.ExampleEnhancedError",
	// 			"caller": "e2h_example_test.go:46",
	// 			"context": "Error executing [bar()] function"
	// 		}
	// 	]
	// }

}
```

## Sample

You can find a sample of the use of go-e2h project [HERE](https://github.com/cdleo/go-e2h/blob/master/e2h_example_test.go)

## Contributing

Comments, suggestions and/or recommendations are always welcomed. Please check the [Contributing Guide](CONTRIBUTING.md) to learn how to get started contributing.
