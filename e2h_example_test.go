/*
Package e2h_test its the test package of the Enhanced Error Handling module
*/
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

	fmt.Printf("Full info JSON (pretty) =>\n\t%s\n", e2h.FormatJSON(err, "github.com", true))

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
	// Full info JSON (pretty) =>
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
