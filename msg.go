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
// In the case of wrapped errors with [fmt.Errorf], Err unwraps the error and renders
// the causal chain as a tree:
//
//	root := errors.New("some deep error")
//	wrapped := fmt.Errorf("could not process file: %w", root)
//	again := fmt.Errorf("failed to do thing: %w", wrapped)
//	msg.Err(again)
//
// Multi-errors produced by [errors.Join] are rendered as branches at the point they
// appear in the chain.
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
// In the case of wrapped errors with [fmt.Errorf], Ferr unwraps the error and renders
// the causal chain as a tree:
//
//	root := errors.New("some deep error")
//	wrapped := fmt.Errorf("could not process file: %w", root)
//	again := fmt.Errorf("failed to do thing: %w", wrapped)
//	msg.Ferr(os.Stderr, again)
//
// Multi-errors produced by [errors.Join] are rendered as branches at the point they
// appear in the chain.
//
// The intended use case for Ferr and [Err] is at the top level of a CLI application where all
// errors are eventually bubbled up to with the appropriate context, which can
// then be shown to end users in a very clear, concise way.
func Ferr(w io.Writer, err error) {
	if err == nil {
		return
	}

	own, children := decompose(err)

	// A bare [errors.Join] at the root has no headline of its own, its message
	// is just its children concatenated. Fall back to a plain render so we don't
	// emit an empty "Error:" line.
	if own == "" && len(children) > 0 {
		Ferror(w, "%v", err)

		return
	}

	Ferror(w, "%v", own)
	renderCauses(w, children, "")
}

// decompose returns err's own message contribution and its child errors. own
// is "" when err is a transparent multi-error such as a bare [errors.Join],
// whose message is exactly its children's joined by "\n" with no headline.
func decompose(err error) (own string, children []error) {
	s := err.Error()

	if multi, ok := err.(interface{ Unwrap() []error }); ok {
		kids := multi.Unwrap()
		joined := errors.Join(kids...).Error()

		if s == joined {
			return "", kids
		}

		if before, found := strings.CutSuffix(s, ": "+joined); found {
			return before, kids
		}

		// Unconventional multi-error format (e.g. [fmt.Errorf] with multiple
		// %w verbs interleaved with text): render as a leaf so children
		// don't appear both inline in the headline and as branches below.
		return s, nil
	}

	if next := errors.Unwrap(err); next != nil {
		if before, found := strings.CutSuffix(s, ": "+next.Error()); found {
			return before, []error{next}
		}

		return s, []error{next}
	}

	return s, nil
}

// renderCauses writes children as a tree under the given line prefix.
func renderCauses(w io.Writer, parents []error, prefix string) {
	causes := decomposeAll(parents)
	for i, c := range causes {
		last := i == len(causes)-1

		connector, nextPrefix := "├─", prefix+"│  "
		if last {
			connector, nextPrefix = "╰─", prefix+"   "
		}

		fmt.Fprintf(w, "%s%s %s: %s\n", prefix, connector, styleCause.Sprint("cause"), c.own)
		renderCauses(w, c.children, nextPrefix)
	}
}

// cause is a renderable tree node: an error decomposed into its own message
// and any further children to recurse into.
type cause struct {
	own      string
	children []error
}

// decomposeAll decomposes each err and inlines transparent multi-errors so
// the returned slice contains only renderable nodes, never an empty
// headline that would print as a stub branch.
func decomposeAll(errs []error) []cause {
	var out []cause

	for _, e := range errs {
		own, children := decompose(e)
		if own == "" && len(children) > 0 {
			out = append(out, decomposeAll(children)...)

			continue
		}

		out = append(out, cause{own: own, children: children})
	}

	return out
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
