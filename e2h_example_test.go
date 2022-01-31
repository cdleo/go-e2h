package e2h_test

import (
	"fmt"

	"github.com/cdleo/go-e2h"
)

func foo() error {
	return fmt.Errorf("foo")
}

func bar() error {
	return e2h.Trace(foo(), "This call wraps the GO error and adds this context and other helpful information")
}

func ExampleEnhancedError() {

	err := e2h.Trace(bar(), "This call adds stack information")

	fmt.Printf("Just Cause => %s\n", e2h.Cause(err))

	fmt.Printf("Just as Error => %s\n", err)

	fmt.Printf("Full info Pretty =>\n%s", e2h.FormatPretty(err, "github.com"))

	fmt.Printf("Full info JSON =>\n\t%s\n", e2h.FormatJSON(err, "github.com"))

	// Output:
	// Just Cause => foo
	// Just as Error => foo: This call wraps the GO error and adds this context and other helpful information
	// Full info Pretty =>
	// - This call adds stack information:
	//   github.com/cdleo/go-e2h_test.ExampleEnhancedError
	//   	github.com/cdleo/go-e2h/e2h_example_test.go:19
	// - This call wraps the GO error and adds this context and other helpful information:
	//   github.com/cdleo/go-e2h_test.bar
	//   	github.com/cdleo/go-e2h/e2h_example_test.go:14
	// - foo
	// Full info JSON =>
	//	{"error":"foo: This call wraps the GO error and adds this context and other helpful information","stack_trace":[{"func":"github.com/cdleo/go-e2h_test.ExampleEnhancedError","caller":"github.com/cdleo/go-e2h/e2h_example_test.go:19","context":"This call adds stack information"},{"func":"github.com/cdleo/go-e2h_test.bar","caller":"github.com/cdleo/go-e2h/e2h_example_test.go:14","context":"This call wraps the GO error and adds this context and other helpful information"}]}

}
