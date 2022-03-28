/*
Package e2hformat is the formatter's package of the Enhanced Error Handling module
*/
package e2hformat

import (
	"fmt"

	"github.com/cdleo/go-commons/formatter"
)

type Format int8

// Allowed output formats.
const (
	Format_Raw Format = iota
	Format_JSON
)

type Params struct {
	//Sets if the output will be beautified
	Beautify bool
	//Sets if at top of the stack shows the last trace (invert = true) or the origin error (invert = false)
	InvertCallstack bool
	//Sets the way in with the filepaths are managed.
	PathHidingMethod formatter.HidingMethod
	//Value to use, according to the selected 'PathHidingMethod'
	PathHidingValue string
}

type Formatter interface {
	Source(err error) string
	Format(err error, params Params) string
}

func NewFormatter(format Format) (Formatter, error) {

	switch format {
	case Format_Raw:
		return newRawFormatter(), nil
	case Format_JSON:
		return newJSONFormatter(), nil
	default:
		return nil, fmt.Errorf("unknown format [%d]", format)
	}

}
