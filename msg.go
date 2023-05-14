// Package msg is a simple, easy to use, opinionated console printing toolkit for Go CLIs rendering pretty
// formatted output with colours specific to the particular message type:
//
//   - Info: For general user information and progress updates
//   - Title: Separation between sections of output
//   - Warn: User warnings
//   - Success: Report success
//   - Error: Report failure
//
// All message types default to stdout other than the `Error` type which prints to stderr by default.
//
// There are also "F-style" print methods that allow you to specify an [io.Writer] to print the messages to.
package msg

import (
	"fmt"
	"io"
	"os"

	"github.com/fatih/color"
)

const (
	// Default colors for each message type.
	colorInfo    = color.FgHiCyan
	colorTitle   = color.FgCyan
	colorWarn    = color.FgYellow
	colorError   = color.FgHiRed
	colorSuccess = color.FgGreen

	// Default statuses for each message type.
	statusInfo    = "Info"
	statusWarn    = "Warning"
	statusError   = "Error"
	statusSuccess = "Success"
)

// Success prints a success message with optional format args to stdout.
//
//	msg.Success("Compiled project: %s", "msg")
func Success(format string, a ...any) {
	Fsuccess(os.Stdout, format, a...)
}

// Fsuccess prints a success message with optional format args to w.
func Fsuccess(w io.Writer, format string, a ...any) {
	success := color.New(colorSuccess, color.Bold)
	message := color.New(color.FgHiWhite)

	fmt.Fprintf(w, "%s: %s\n", success.Sprint(statusSuccess), message.Sprintf(format, a...))
}

// Error prints an error message with optional format args to stderr.
//
//	msg.Error("Invalid config")
//	msg.Error("Could not find file: %s", "missing.txt")
func Error(format string, a ...any) {
	Ferror(os.Stderr, format, a...)
}

// Ferror prints an error message with optional format args to w.
func Ferror(w io.Writer, format string, a ...any) {
	e := color.New(colorError, color.Bold)
	message := color.New(color.FgHiWhite)

	fmt.Fprintf(w, "%s: %s\n", e.Sprint(statusError), message.Sprintf(format, a...))
}

// Warn prints a warning message with optional format args to stdout.
//
//	msg.Warn("Skipping %s, directory is empty", "some/empty/dir")
func Warn(format string, a ...any) {
	Fwarn(os.Stdout, format, a...)
}

// Fwarn prints a warning message with optional format args to w.
func Fwarn(w io.Writer, format string, a ...any) {
	warn := color.New(colorWarn, color.Bold)
	message := color.New(color.FgHiWhite)

	fmt.Fprintf(w, "%s: %s\n", warn.Sprint(statusWarn), message.Sprintf(format, a...))
}
