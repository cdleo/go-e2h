/*
Package e2hformat is the formatter's package of the Enhanced Error Handling module
*/
package e2hformat

import (
	"fmt"
	"path/filepath"
	"strings"
)

type Format int8

// Allowed output formats.
const (
	Format_Raw Format = iota
	Format_JSON
)

type HidingMethod int8

//Allowed path treatments
const (
	HidingMethod_None HidingMethod = iota

	HidingMethod_FullBaseline

	HidingMethod_ToFolder
)

type Params struct {
	//Sets if the output will be beautified
	Beautify bool
	//Sets if at top of the stack shows the last trace (invert = true) or the origin error (invert = false)
	InvertCallstack bool
	//Sets the way in with the filepaths are managed.
	PathHidingMethod HidingMethod
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

	file = filepath.Clean(file)
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
