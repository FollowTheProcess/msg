/*
Scratchpad for development. msg is a library.
*/

package main

import (
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

	printer.Info("Something happened")

	printer.SymbolInfo = "🔎"

	printer.Info("Something happened with a different symbol!")
}
