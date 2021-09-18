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

	name := "Tom"

	printer.Title("Some stuff below")

	msg.Warnf("Warning you about %s", name)
	msg.Textf("I'm some text below the warning about %s", name)
}
