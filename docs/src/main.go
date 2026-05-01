package main

import (
	"errors"
	"fmt"

	"go.followtheprocess.codes/msg"
)

func main() {
	msg.Title("Your CLI output:")
	msg.Success("compiled 42 packages")
	msg.Warn("directory is empty, skipping")
	msg.Info("using value from config file")

	fmt.Println()

	// Wrapped chains and errors.Join values render as a tree, with branches
	// for every parallel cause and vertical continuation through deeper levels.
	sshIssue := fmt.Errorf("ssh key %q: %w", "id_rsa", errors.New("file not found"))
	dbIssue := fmt.Errorf(
		"database: %w",
		fmt.Errorf("dial tcp 10.0.0.5:5432: %w", errors.New("connection refused")),
	)
	configIssue := errors.New("missing required field 'region'")

	checks := errors.Join(sshIssue, dbIssue, configIssue)
	err := fmt.Errorf("preflight checks failed: %w", checks)

	msg.Err(err)
}
