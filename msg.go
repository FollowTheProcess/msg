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
package msg // import "go.followtheprocess.codes/msg"

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"go.followtheprocess.codes/hue"
)

const (
	styleError   = hue.Red | hue.Bold
	styleTitle   = hue.BrightCyan
	styleInfo    = hue.Cyan | hue.Bold
	styleWarn    = hue.Yellow | hue.Bold
	styleSuccess = hue.Green | hue.Bold
	styleCause   = hue.Bold
)

const (
	// Default statuses for each message type.
	statusInfo    = "Info"
	statusWarn    = "Warning"
	statusError   = "Error"
	statusSuccess = "Success"
)

// ColorEnabled sets whether the output from this package is colourised.
//
// msg defaults to automatic detection based on a number of attributes:
//   - The value of $NO_COLOR and/or $FORCE_COLOR
//   - The value of $TERM (xterm enables colour)
//   - Whether [os.Stdout] is pointing to a terminal
//
// This means that msg should do a reasonable job of auto-detecting when to colourise output
// and should not write escape sequences when piping between processes or when writing to files etc.
//
// This function may be called to bypass the above detection and explicitly set the value, useful in CLI
// applications where a --no-color flag might be expected.
//
// ColorEnabled may be called safely from concurrently executing goroutines.
func ColorEnabled(v bool) {
	hue.Enabled(v)
}

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

// Err prints a nicely formatted error message from an actual error to [os.Stderr], bypassing
// the need for the caller to construct the format string.
//
//	err := errors.New("Uh oh!")
//	msg.Err(err) // Equivalent to msg.Error("%v", err)
//
// In the case of wrapped errors with [fmt.Errorf], Err recursively unwraps the error,
// showing each of the errors in the causal chain in a tree-like structure.
//
//	root := errors.New("some deep error")
//	wrapped := fmt.Errorf("could not process file: %w", root)
//	again := fmt.Errorf("failed to do thing: %w", wrapped)
//	msg.Err(again) // Unwraps the above and shows each cause as a new indented line
//
// The intended use case for Err is at the top level of a CLI application where all
// errors are eventually bubbled up to with the appropriate context, which Err can
// then show to end users in a very clear, concise way.
func Err(err error) {
	Ferr(os.Stderr, err)
}

// Ferror prints an error message with optional format args to w.
//
//	msg.Ferror(os.Stderr, "Uh oh! %s", "something wrong")
func Ferror(w io.Writer, format string, a ...any) {
	fmt.Fprintf(w, "%s: %s\n", styleError.Sprint(statusError), fmt.Sprintf(format, a...))
}

// Ferr prints a nicely formatted error message from an actual error to w, bypassing
// the need for the caller to construct the format string.
//
//	err := errors.New("Uh oh!")
//	msg.Ferr(os.Stderr, err) // Equivalent to msg.Err(err)
//
// In the case of wrapped errors with [fmt.Errorf], Ferr recursively unwraps the error,
// showing each of the errors in the causal chain in a tree-like structure.
//
//	root := errors.New("some deep error")
//	wrapped := fmt.Errorf("could not process file: %w", root)
//	again := fmt.Errorf("failed to do thing: %w", wrapped)
//	msg.Ferr(os.Stderr, again) // Unwraps the above and shows each cause as a new indented line
//
// The intended use case for Ferr and [Err] is at the top level of a CLI application where all
// errors are eventually bubbled up to with the appropriate context, which can
// then be shown to end users in a very clear, concise way.
func Ferr(w io.Writer, err error) {
	if err == nil {
		return
	}

	// No wrapped errors, just do what [Error] does
	if errors.Unwrap(err) == nil {
		Ferror(w, "%v", err)
		return
	}

	// TODO(@FollowTheProcess): We should be able to build this stack of errors by
	// calling Unwrap alone as that is less fragile, but because each layer of unwrap contains
	// all the child elements too we need something a bit clever to recurse all the way down to <nil>, then
	// build the stack back up from the bottom up. For example you'd get something like:
	// failed to do something: could not find file: invalid permissions: super deep error
	// cause: could not find file: invalid permissions: super deep error
	// cause: invalid permissions: super deep error
	// cause: super deep error
	//
	// With each level having all the child errors in so you get lots of duplication, splitting
	// on colons is a bit of a hack that relies on convention `fmt.Errorf("some error: %w", err)`
	// but it works well enough for me for now as I always do that anyway

	chain := strings.Split(err.Error(), ":")
	root := chain[0]
	causes := chain[1:]

	Ferror(w, "%v", strings.TrimSpace(root))

	indent := 0
	for _, cause := range causes {
		fmt.Fprintf(w, "%s╰─ %s: %v\n", strings.Repeat(" ", indent), styleCause.Sprint("cause"), strings.TrimSpace(cause))
		indent += 3
	}
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
