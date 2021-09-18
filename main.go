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

	printer.Title("I'm a title")
	fmt.Println("I'm some text on the next line")

	s := printer.TitleString("I'm a titlestring")
	fmt.Println(s)

	printer.SymbolTitle = "💨"

	s = printer.TitleString("I'm a titlestring with a symbol")
	fmt.Println(s)

	// If you just want to use the default
	msg.Title("I'm the default title")
}
