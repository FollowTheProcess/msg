/*
Scratchpad for development. msg is a library.
*/

package main

import (
	"fmt"
	"os"

	"github.com/FollowTheProcess/msg/msg"
	"github.com/fatih/color"
)

func main() {
	printer := msg.Printer{
		SymbolInfo:  "ℹ",
		SymbolTitle: "",
		SymbolWarn:  "⚠️",
		SymbolFail:  "✘",
		SymbolGood:  "✔",
		ColorInfo:   color.FgCyan,
		ColorTitle:  color.FgCyan,
		ColorWarn:   color.FgYellow,
		ColorFail:   color.FgRed,
		ColorGood:   color.FgGreen,
		Out:         os.Stdout,
	}

	msg.Title("Stuff below here")

	printer.Fail("I'm an error")
	fmt.Println("I'm below the error")

	s := printer.FailString("I'm an error string")
	fmt.Println(s)

	printer.SymbolFail = "🤬"
	printer.Fail("I'm an error with a different symbol")

	msg.Fail("I'm the default error")
}
