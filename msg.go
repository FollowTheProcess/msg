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

	"github.com/FollowTheProcess/msg/internal/colour"
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
	fmt.Fprintf(w, "%s: %s\n", colour.Success(statusSuccess), fmt.Sprintf(format, a...))
}

// Error prints an error message with optional format args to stderr.
//
//	msg.Error("Invalid config")
//	msg.Error("Could not find file: %s", "missing.txt")
func Error(format string, a ...any) {
	Ferror(os.Stderr, format, a...)
}

// Ferror prints an error message with optional format args to w.
//
//	msg.Ferror(os.Stderr, "Uh oh! %s", "something wrong")
func Ferror(w io.Writer, format string, a ...any) {
	fmt.Fprintf(w, "%s: %s\n", colour.Error(statusError), fmt.Sprintf(format, a...))
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
	fmt.Fprintf(w, "%s: %s\n", colour.Warn(statusWarn), fmt.Sprintf(format, a...))
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
	fmt.Fprintf(w, "%s: %s\n", colour.Info(statusInfo), fmt.Sprintf(format, a...))
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
	fmt.Fprintf(w, "\n%s\n\n", colour.Title(fmt.Sprintf(format, a...)))
}
