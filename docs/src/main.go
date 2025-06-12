package main

import (
	"errors"
	"fmt"

	"followtheprocess.codes/msg"
)

func main() {
	msg.Title("Your CLI output:")
	msg.Success("compiled 42 packages")
	msg.Warn("directory is empty, skipping")
	msg.Info("using value from config file")

	fmt.Println()

	// Simulate a wrapped tree of errors throughout
	// your application
	err := errors.New("bad file permissions")
	one := fmt.Errorf("could not read DB config: %w", err)
	two := fmt.Errorf("failed to insert new record: %w", one)
	three := fmt.Errorf("could not complete transaction: %w", two)

	// Bubble them all up
	msg.Err(three)
}
