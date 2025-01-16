package msg_test

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/FollowTheProcess/hue"
	"github.com/FollowTheProcess/msg"
)

const (
	successCode = "\x1b[1;32m"
	errorCode   = "\x1b[1;31m"
	causeCode   = "\x1b[1m"
	warnCode    = "\x1b[1;33m"
	infoCode    = "\x1b[1;36m"
	titleCode   = "\x1b[96m"
	resetCode   = "\x1b[0m"
)

func TestMain(m *testing.M) {
	// Disable colour auto-detection so CI passes
	hue.Enabled(true)
	code := m.Run()
	os.Exit(code)
}

func TestSuccess(t *testing.T) {
	buf := new(bytes.Buffer)
	msg.Fsuccess(buf, "Something went well: %v", 42)

	want := fmt.Sprintf("%sSuccess%s: Something went well: 42\n", successCode, resetCode)

	if got := buf.String(); got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestSuccessCaptured(t *testing.T) {
	successFunc := func() {
		msg.Success("Worked")
	}
	got := captureStdout(t, successFunc)
	want := fmt.Sprintf("%sSuccess%s: Worked\n", successCode, resetCode)

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestError(t *testing.T) {
	buf := new(bytes.Buffer)
	msg.Ferror(buf, "Something broke: %v", true)

	want := fmt.Sprintf("%sError%s: Something broke: true\n", errorCode, resetCode)

	if got := buf.String(); got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestErrorCaptured(t *testing.T) {
	errorFunc := func() {
		msg.Error("Bad number (%v)", 42)
	}
	got := captureStderr(t, errorFunc)
	want := fmt.Sprintf("%sError%s: Bad number (42)\n", errorCode, resetCode)

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestErr(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		buf := new(bytes.Buffer)
		var err error = nil
		msg.Ferr(buf, err)

		if buf.Len() != 0 {
			t.Fatalf("nil err wrote output: %s", buf.String())
		}
	})
	t.Run("plain", func(t *testing.T) {
		buf := new(bytes.Buffer)
		err := errors.New("Something broke")
		msg.Ferr(buf, err)

		want := fmt.Sprintf("%sError%s: Something broke\n", errorCode, resetCode)

		if got := buf.String(); got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	})

	t.Run("wrapped", func(t *testing.T) {
		buf := new(bytes.Buffer)
		root := errors.New("something broke")
		one := fmt.Errorf("could not frobnicate the baz: %w", root)
		two := fmt.Errorf("failed to process file: %w", one)
		three := fmt.Errorf("could not deposit money: %w", two)
		msg.Ferr(buf, three)
		wantTemplate := `%[1]sError%[2]s: could not deposit money
╰─ %[3]scause%[2]s: failed to process file
   ╰─ %[3]scause%[2]s: could not frobnicate the baz
      ╰─ %[3]scause%[2]s: something broke
`

		want := fmt.Sprintf(wantTemplate, errorCode, resetCode, causeCode)

		if got := buf.String(); got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	})
}

func TestErrCaptured(t *testing.T) {
	t.Run("plain", func(t *testing.T) {
		errorFunc := func() {
			err := errors.New("Bang!")
			msg.Err(err)
		}
		got := captureStderr(t, errorFunc)
		want := fmt.Sprintf("%sError%s: Bang!\n", errorCode, resetCode)

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	})
	t.Run("wrapped", func(t *testing.T) {
		errorFunc := func() {
			root := errors.New("bang!")
			one := fmt.Errorf("dingle cannot dangle on version <2: %w", root)
			two := fmt.Errorf("could not read file: %w", one)
			three := fmt.Errorf("you have no money: %w", two)
			msg.Err(three)
		}
		wantTemplate := `%[1]sError%[2]s: you have no money
╰─ %[3]scause%[2]s: could not read file
   ╰─ %[3]scause%[2]s: dingle cannot dangle on version <2
      ╰─ %[3]scause%[2]s: bang!
`
		got := captureStderr(t, errorFunc)
		want := fmt.Sprintf(wantTemplate, errorCode, resetCode, causeCode)

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	})
}

func TestWarn(t *testing.T) {
	buf := new(bytes.Buffer)
	msg.Fwarn(buf, "skipping directory %s", "./tmp")

	want := fmt.Sprintf("%sWarning%s: skipping directory ./tmp\n", warnCode, resetCode)

	if got := buf.String(); got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestWarnCaptured(t *testing.T) {
	warnFunc := func() {
		msg.Warn("Skipping something (%d)", 42)
	}
	got := captureStdout(t, warnFunc)
	want := fmt.Sprintf("%sWarning%s: Skipping something (42)\n", warnCode, resetCode)

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestInfo(t *testing.T) {
	buf := new(bytes.Buffer)
	msg.Finfo(buf, "You have %d projects on GitHub", 27)

	want := fmt.Sprintf("%sInfo%s: You have 27 projects on GitHub\n", infoCode, resetCode)

	if got := buf.String(); got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestInfoCaptured(t *testing.T) {
	infoFunc := func() {
		msg.Info("You are %d years old", 29)
	}
	got := captureStdout(t, infoFunc)
	want := fmt.Sprintf("%sInfo%s: You are 29 years old\n", infoCode, resetCode)

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestTitle(t *testing.T) {
	buf := new(bytes.Buffer)
	msg.Ftitle(buf, "Section Header")

	want := fmt.Sprintf("\n%sSection Header%s\n\n", titleCode, resetCode)

	if got := buf.String(); got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestTitleCaptured(t *testing.T) {
	titleFunc := func() {
		msg.Title("Section Header")
	}
	got := captureStdout(t, titleFunc)
	want := fmt.Sprintf("\n%sSection Header%s\n\n", titleCode, resetCode)

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestVisual(t *testing.T) {
	msg.Title("Your CLI output:")
	msg.Success("compiled 42 packages")
	msg.Warn("directory is empty, skipping")
	msg.Info("using value from config file")

	fmt.Println()

	err := errors.New("bad file permissions")
	one := fmt.Errorf("could not read DB config: %w", err)
	two := fmt.Errorf("failed to insert new record: %w", one)
	three := fmt.Errorf("could not complete transaction: %w", two)

	msg.Err(three)
}

func captureStdout(t *testing.T, printer func()) string {
	t.Helper()
	old := os.Stdout // Backup of the real one
	defer func() {
		os.Stdout = old // Set it back even if we error later
	}()

	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("os.Pipe() returned an error: %v", err)
	}

	// Set stdout to our new pipe
	os.Stdout = w

	capture := make(chan string)
	// Copy in a goroutine so printing can't block forever
	go func() {
		buf := new(bytes.Buffer)
		io.Copy(buf, r) //nolint: errcheck
		capture <- buf.String()
	}()

	// Call our test function that prints to stdout
	printer()

	// Close the writer
	w.Close()
	captured := <-capture

	return captured
}

func captureStderr(t *testing.T, printer func()) string {
	t.Helper()
	old := os.Stderr // Backup of the real one
	defer func() {
		os.Stderr = old // Set it back even if we error later
	}()

	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("os.Pipe() returned an error: %v", err)
	}

	// Set stderr to our new pipe
	os.Stderr = w

	capture := make(chan string)
	// Copy in a goroutine so printing can't block forever
	go func() {
		buf := new(bytes.Buffer)
		io.Copy(buf, r) //nolint: errcheck
		capture <- buf.String()
	}()

	// Call our test function that prints to stderr
	printer()

	// Close the writer
	w.Close()
	captured := <-capture

	return captured
}
