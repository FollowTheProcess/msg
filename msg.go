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

	"github.com/FollowTheProcess/hue"
)

const (
	styleError   = hue.Red | hue.Bold
	styleTitle   = hue.BrightCyan
	styleInfo    = hue.Cyan | hue.Bold
	styleWarn    = hue.Yellow | hue.Bold
	styleSuccess = hue.Green | hue.Bold
)

const (
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
//
//	msg.Fsuccess(os.Stdout, "Compiled project: %s", "msg")
func Fsuccess(w io.Writer, format string, a ...any) {
	fmt.Fprintf(w, "%s: %s\n", styleSuccess.Sprint(statusSuccess), fmt.Sprintf(format, a...))
}

// Error prints an error message with optional format args to stderr.
//
//	msg.Error("Invalid config")
//	msg.Error("Could not find file: %s", "missing.txt")
func Error(format string, a ...any) {
	Ferror(os.Stderr, format, a...)
}

// Err is a convenience wrapper around [Error] allowing passing an error directly
// without the need for the %v format verb.
//
//	err := errors.New("Uh oh!")
//	msg.Err(err) // Equivalent to msg.Error("%v", err)
func Err(err error) {
	Error("%v", err)
}

// Ferror prints an error message with optional format args to w.
//
//	msg.Ferror(os.Stderr, "Uh oh! %s", "something wrong")
func Ferror(w io.Writer, format string, a ...any) {
	fmt.Fprintf(w, "%s: %s\n", styleError.Sprint(statusError), fmt.Sprintf(format, a...))
}

// Ferr is a convenience wrapper around [Ferror] allowing passing an error directly
// without the new for the %v format verb.
//
// It is the 'F' equivalent of [Err], taking an [io.Writer] to print the error to.
//
//	err := errors.New("Uh oh!")
//	msg.Ferr(os.Stderr, err) // Equivalent to msg.Ferror(os.Stderr, "%v", err)
func Ferr(w io.Writer, err error) {
	Ferror(w, "%v", err)
}

// Warn prints a warning message with optional format args to stdout.
//
//	msg.Warn("Skipping %s, directory is empty", "some/empty/dir")
func Warn(format string, a ...any) {
	Fwarn(os.Stdout, format, a...)
}

// Fwarn prints a warning message with optional format args to w.
//
//	msg.Fwarn(os.Stderr, "hmmmm: %v", true)
func Fwarn(w io.Writer, format string, a ...any) {
	fmt.Fprintf(w, "%s: %s\n", styleWarn.Sprint(statusWarn), fmt.Sprintf(format, a...))
}

// Info prints an info message with optional format args to stdout.
//
//	msg.Info("You have %d repos on GitHub", 42)
func Info(format string, a ...any) {
	Finfo(os.Stdout, format, a...)
}

// Finfo prints an info message with optional format args to w.
//
//	msg.Finfo(os.Stdout, "The meaning of life is %v", 42)
func Finfo(w io.Writer, format string, a ...any) {
	fmt.Fprintf(w, "%s: %s\n", styleInfo.Sprint(statusInfo), fmt.Sprintf(format, a...))
}

// Title prints a title message to stdout.
//
// A title message differs from every other message type in msg as it
// has 1 leading newline and 2 trailing newlines to create separation between
// the sections it is differentiating in your CLI.
//
//	msg.Title("Some section")
func Title(format string, a ...any) {
	Ftitle(os.Stdout, format, a...)
}

// Ftitle prints a title message to w.
//
// A title message differs from every other message type in msg as it
// has 1 leading newline and 2 trailing newlines to create separation between
// the sections it is differentiating in your CLI.
//
//	msg.Ftitle(os.Stdout, "Some section")
func Ftitle(w io.Writer, format string, a ...any) {
	fmt.Fprintf(w, "\n%s\n\n", styleTitle.Sprint(fmt.Sprintf(format, a...)))
}
