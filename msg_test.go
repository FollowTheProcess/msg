package msg_test

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/FollowTheProcess/msg"
	"github.com/FollowTheProcess/msg/internal/colour"
)

func TestSuccess(t *testing.T) {
	buf := new(bytes.Buffer)
	msg.Fsuccess(buf, "Something went well: %v", 42)

	want := fmt.Sprintf("%sSuccess%s: Something went well: 42\n", colour.CodeSuccess, colour.CodeReset)

	if got := buf.String(); got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestSuccessCaptured(t *testing.T) {
	successFunc := func() {
		msg.Success("Worked")
	}
	got := captureStdout(t, successFunc)
	want := fmt.Sprintf("%sSuccess%s: Worked\n", colour.CodeSuccess, colour.CodeReset)

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestError(t *testing.T) {
	buf := new(bytes.Buffer)
	msg.Ferror(buf, "Something broke: %v", true)

	want := fmt.Sprintf("%sError%s: Something broke: true\n", colour.CodeError, colour.CodeReset)

	if got := buf.String(); got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestErrorCaptured(t *testing.T) {
	errorFunc := func() {
		msg.Error("Bad number (%v)", 42)
	}
	got := captureStderr(t, errorFunc)
	want := fmt.Sprintf("%sError%s: Bad number (42)\n", colour.CodeError, colour.CodeReset)

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestErr(t *testing.T) {
	buf := new(bytes.Buffer)
	err := errors.New("Something broke")
	msg.Ferr(buf, err)

	want := fmt.Sprintf("%sError%s: Something broke\n", colour.CodeError, colour.CodeReset)

	if got := buf.String(); got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestErrCaptured(t *testing.T) {
	errorFunc := func() {
		err := errors.New("Bang!")
		msg.Err(err)
	}
	got := captureStderr(t, errorFunc)
	want := fmt.Sprintf("%sError%s: Bang!\n", colour.CodeError, colour.CodeReset)

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestWarn(t *testing.T) {
	buf := new(bytes.Buffer)
	msg.Fwarn(buf, "skipping directory %s", "./tmp")

	want := fmt.Sprintf("%sWarning%s: skipping directory ./tmp\n", colour.CodeWarn, colour.CodeReset)

	if got := buf.String(); got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestWarnCaptured(t *testing.T) {
	warnFunc := func() {
		msg.Warn("Skipping something (%d)", 42)
	}
	got := captureStdout(t, warnFunc)
	want := fmt.Sprintf("%sWarning%s: Skipping something (42)\n", colour.CodeWarn, colour.CodeReset)

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestInfo(t *testing.T) {
	buf := new(bytes.Buffer)
	msg.Finfo(buf, "You have %d projects on GitHub", 27)

	want := fmt.Sprintf("%sInfo%s: You have 27 projects on GitHub\n", colour.CodeInfo, colour.CodeReset)

	if got := buf.String(); got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestInfoCaptured(t *testing.T) {
	infoFunc := func() {
		msg.Info("You are %d years old", 29)
	}
	got := captureStdout(t, infoFunc)
	want := fmt.Sprintf("%sInfo%s: You are 29 years old\n", colour.CodeInfo, colour.CodeReset)

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestTitle(t *testing.T) {
	buf := new(bytes.Buffer)
	msg.Ftitle(buf, "Section Header")

	want := fmt.Sprintf("\n%sSection Header%s\n\n", colour.CodeTitle, colour.CodeReset)

	if got := buf.String(); got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestTitleCaptured(t *testing.T) {
	titleFunc := func() {
		msg.Title("Section Header")
	}
	got := captureStdout(t, titleFunc)
	want := fmt.Sprintf("\n%sSection Header%s\n\n", colour.CodeTitle, colour.CodeReset)

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
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
