# GO-E2H

GO Enhanced Error Handling (a.k.a. go-e2h) is a lightweight Golang module to add a better stack trace and context information on error events.

## Usage

We use an object of the provided interface EnhancedError to store the context and stack information:

```go
type EnhancedError interface {
    // This function returns the Error string plus the origin custom message (if exists)
	Error() string
    // This function returns the source error
	Cause() error
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

Additionally, this module provide functions to pretty print the error information, over different outputs:

```go
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

	"github.com/cdleo/go-e2h"
)

func foo() error {
	return fmt.Errorf("foo")
}

func bar() error {
	return e2h.Tracem(foo(), "This call wraps the GO standard error and adds this message and other helpful information")
}

func ExampleEnhancedError() {

	err := e2h.Tracef(bar(), "This call adds this %s message and stack information", "formatted")

	fmt.Printf("Just cause => %s\n", e2h.Source(err))

	fmt.Printf("Just as Error => %v\n", err)

	fmt.Printf("Full info (pretty) =>\n%s", e2h.FormatPretty(err, "github.com"))

	fmt.Printf("Full info JSON (indented) =>\n\t%s\n", e2h.FormatJSON(err, "github.com", true))

	fmt.Printf("Full info JSON (std) =>\n\t%s\n", e2h.FormatJSON(err, "github.com", false))

	// Output:
	// Just cause => foo
	// Just as Error => foo: This call wraps the GO standard error and adds this message and other helpful information
	// Full info (pretty) =>
	// - This call adds this formatted message and stack information:
	//   github.com/cdleo/go-e2h_test.ExampleEnhancedError
	//   	github.com/cdleo/go-e2h/e2h_example_test.go:22
	// - This call wraps the GO standard error and adds this message and other helpful information:
	//   github.com/cdleo/go-e2h_test.bar
	//   	github.com/cdleo/go-e2h/e2h_example_test.go:17
	// - foo
	// Full info JSON (indented) =>
	// 	{
	// 	"error": "foo: This call wraps the GO standard error and adds this message and other helpful information",
	// 	"stack_trace": [
	// 		{
	// 			"func": "github.com/cdleo/go-e2h_test.ExampleEnhancedError",
	// 			"caller": "github.com/cdleo/go-e2h/e2h_example_test.go:22",
	// 			"context": "This call adds this formatted message and stack information"
	// 		},
	// 		{
	// 			"func": "github.com/cdleo/go-e2h_test.bar",
	// 			"caller": "github.com/cdleo/go-e2h/e2h_example_test.go:17",
	// 			"context": "This call wraps the GO standard error and adds this message and other helpful information"
	// 		}
	// 	]
	// }
	// Full info JSON (std) =>
	// 	{"error":"foo: This call wraps the GO standard error and adds this message and other helpful information","stack_trace":[{"func":"github.com/cdleo/go-e2h_test.ExampleEnhancedError","caller":"github.com/cdleo/go-e2h/e2h_example_test.go:22","context":"This call adds this formatted message and stack information"},{"func":"github.com/cdleo/go-e2h_test.bar","caller":"github.com/cdleo/go-e2h/e2h_example_test.go:17","context":"This call wraps the GO standard error and adds this message and other helpful information"}]}

}
```

## Sample

You can find a sample of the use of go-e2h project [HERE](https://github.com/cdleo/go-e2h/blob/master/e2h_example_test.go)

## Contributing

Comments, suggestions and/or recommendations are always welcomed. Please check the [Contributing Guide](CONTRIBUTING.md) to learn how to get started contributing.
