package main

import (
	"github.com/FollowTheProcess/msg"
)

func main() {
	msg.Title("Your CLI Output:")
	msg.Success("compiled your project")
	msg.Warn("directory is empty")
	msg.Info("using value from config file")
	msg.Error("file not found %s", "missing.txt")
}
