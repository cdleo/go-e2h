/*
Package e2hformat is the formatter's package of the Enhanced Error Handling module
*/
package e2hformat

import (
	"fmt"
	"strings"
)

// Level defines log levels.
type Format int8
type HidingMethod int8

const (
	// Disabled disables the logger.
	Format_Raw Format = iota

	// Mensajes de muy baja frecuencia que se deben mostrar siempre (como el copyright)
	Format_JSON
)

const (
	HidingMethod_None HidingMethod = iota

	HidingMethod_FullBaseline

	HidingMethod_ToFolder
)

type Params struct {
	Beautify         bool
	InvertCallstack  bool
	PathHidingMethod HidingMethod
	PathHidingValue  string
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

// This function format the sourcefile according to the provided params
func formatSourceFile(file string, hidingMethod HidingMethod, hidingValue string) string {
	switch hidingMethod {
	case HidingMethod_FullBaseline:
		return removePathSegment(file, hidingValue)
	case HidingMethod_ToFolder:
		return removeBeforeFolder(file, hidingValue)
	default: //HidingMethod_None
		return file
	}
}

// Utility funtion that removes the first part of the filepath til the end of `baseline` path argument
func removePathSegment(file string, baseline string) string {

	prettyCaller := strings.ReplaceAll(file, baseline, "")
	if len(prettyCaller) > 0 {
		return prettyCaller
	}

	return file
}

// Utility funtion that removes the first part of the filepath til found the folder indicated in `newRootFolder` argument
func removeBeforeFolder(file string, newRootFolder string) string {

	fileParts := strings.Split(file, newRootFolder)
	if len(fileParts) < 2 {
		return file
	}
	return newRootFolder + fileParts[1]
}
