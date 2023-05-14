package msg_test

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/FollowTheProcess/msg"
)

func TestSuccess(t *testing.T) {
	buf := new(bytes.Buffer)
	msg.Fsuccess(buf, "Something went well: %v", 42)

	want := "Success: Something went well: 42\n"

	if got := buf.String(); got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestSuccessCaptured(t *testing.T) {
	successFunc := func() {
		msg.Success("Worked")
	}
	got := captureStdout(t, successFunc)
	want := "Success: Worked\n"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestError(t *testing.T) {
	buf := new(bytes.Buffer)
	msg.Ferror(buf, "Something broke: %v", true)

	want := "Error: Something broke: true\n"

	if got := buf.String(); got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestErrorCaptured(t *testing.T) {
	errorFunc := func() {
		msg.Error("Bad number (%v)", 42)
	}
	got := captureStderr(t, errorFunc)
	want := "Error: Bad number (42)\n"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestWarn(t *testing.T) {
	buf := new(bytes.Buffer)
	msg.Fwarn(buf, "skipping directory %s", "./tmp")

	want := "Warning: skipping directory ./tmp\n"

	if got := buf.String(); got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestWarnCaptured(t *testing.T) {
	warnFunc := func() {
		msg.Warn("Skipping something (%d)", 42)
	}
	got := captureStdout(t, warnFunc)
	want := "Warning: Skipping something (42)\n"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestInfo(t *testing.T) {
	buf := new(bytes.Buffer)
	msg.Finfo(buf, "You have %d projects on GitHub", 27)

	want := "Info: You have 27 projects on GitHub\n"

	if got := buf.String(); got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestInfoCaptured(t *testing.T) {
	infoFunc := func() {
		msg.Info("You are %d years old", 29)
	}
	got := captureStdout(t, infoFunc)
	want := "Info: You are 29 years old\n"

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
