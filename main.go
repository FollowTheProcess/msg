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

	printer.Warn("I'm a warning")
	fmt.Println("I'm under the warning")

	// Remove the symbol
	printer.SymbolWarn = ""
	printer.Warn("I'm a warning without a symbol")
	fmt.Println("I'm under this warning")

	printer.SymbolWarn = "⛔️"
	printer.Warn("I'm a warning with this symbol now")
	fmt.Println("I'm under yet another warning")

	msg.Warn("Default warning")
}
