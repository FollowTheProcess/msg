// Package colour implements basic text colouring for msg's limited needs.
//
// In particular, it's not expected to provide every ANSI code, just the ones we need. The codes have also been padded so that they are
// the same length, which means [text/tabwriter] will correctly calculate alignment as long as styles are not mixed within a table.
//
// The functions in this package respect both $NO_COLOR and $FORCE_COLOR.
package colour

import (
	"os"
	"sync"
)

// ANSI codes for coloured output, they are all the same length so as not to throw off
// alignment of [text/tabwriter].
const (
	CodeError   = "\x1b[1;0031m" // Bold red, used for the error status
	CodeTitle   = "\x1b[1;0096m" // Bold hi cyan, used for titles
	CodeInfo    = "\x1b[1;0036m" // Bold cyan, used for info
	CodeWarn    = "\x1b[1;0033m" // Bold yellow, used for warn
	CodeSuccess = "\x1b[1;0032m" // Bold geen, used for success
	CodeReset   = "\x1b[000000m" // Reset all attributes
)

// getColourOnce is a [sync.OnceValues] function that returns the state of
// $NO_COLOR and $FORCE_COLOR, once and only once to avoid us calling
// os.Getenv on every call to a colour function.
var getColourOnce = sync.OnceValues(getColour)

// getColour returns whether $NO_COLOR and $FORCE_COLOR were set.
func getColour() (noColour bool, forceColour bool) {
	no := os.Getenv("NO_COLOR") != ""
	force := os.Getenv("FORCE_COLOR") != ""

	return no, force
}

// Title returns a title styled string.
func Title(text string) string {
	return sprint(CodeTitle, text)
}

// Success returns a success styled string.
func Success(text string) string {
	return sprint(CodeSuccess, text)
}

// Info returns an info styled string.
func Info(text string) string {
	return sprint(CodeInfo, text)
}

// Warn returns a warn styled string.
func Warn(text string) string {
	return sprint(CodeWarn, text)
}

// Error returns an error styled string.
func Error(text string) string {
	return sprint(CodeError, text)
}

// sprint returns a string with a given colour and the reset code.
//
// It handles checking for NO_COLOR and FORCE_COLOR.
func sprint(code, text string) string {
	no, force := getColourOnce()

	// $FORCE_COLOR overrides $NO_COLOR
	if force {
		return code + text + CodeReset
	}

	// $NO_COLOR is next
	if no {
		return text
	}

	// Normal
	return code + text + CodeReset
}
