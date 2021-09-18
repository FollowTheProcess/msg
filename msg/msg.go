package msg

import (
	"fmt"
	"io"
	"os"

	"github.com/fatih/color"
)

const (
	// Default symbols
	defaultInfoSymbol  = "ℹ"
	defaultTitleSymbol = ""
	defaultWarnSymbol  = "⚠️"
	defaultFailSymbol  = "✘"
	defaultGoodSymbol  = "✔"

	// Default colors
	defaultInfoColor  = color.FgCyan
	defaultTitleColor = color.FgCyan
	defaultWarnColor  = color.FgYellow
	defaultFailColor  = color.FgRed
	defaultGoodColor  = color.FgGreen
)

// Printer is the primary construct in msg, it allows you to configure the colors
// and symbols used for each of the printing methods attached to it.
type Printer struct {
	// Symbols

	// The symbol used to prefix Info and InfoString
	SymbolInfo string
	// The symbol used to prefix Title and TitleString
	SymbolTitle string
	// The symbol used ot prefix Warn and WarnString
	SymbolWarn string
	// The symbol used to prefix Fail and FailString
	SymbolFail string
	// The symbol used to prefix Good and GoodString
	SymbolGood string

	// Colors

	// The color used for Info and InfoString
	ColorInfo color.Attribute
	// The color used for Title and TitleString
	ColorTitle color.Attribute
	// The color used for Warn and WarnString
	ColorWarn color.Attribute
	// The color used for Fail and FailString
	ColorFail color.Attribute
	// The color used for Good and GoodString
	ColorGood color.Attribute

	// Output

	// Where to do the printing, useful for testing
	Out io.Writer
}

// Default constructs and returns a default Printer with sensible colors and symbols
// if you just want a nice, standard method of printing things to the user and don't want
// to customise anything, use the returned Printer from this function.
func Default() Printer {
	return Printer{
		SymbolInfo:  defaultInfoSymbol,
		SymbolTitle: defaultTitleSymbol,
		SymbolWarn:  defaultWarnSymbol,
		SymbolFail:  defaultFailSymbol,
		SymbolGood:  defaultGoodSymbol,
		ColorInfo:   defaultInfoColor,
		ColorTitle:  defaultTitleColor,
		ColorWarn:   defaultWarnColor,
		ColorFail:   defaultFailColor,
		ColorGood:   defaultGoodColor,
		Out:         os.Stdout,
	}
}

func (p *Printer) Title(text string) {
	title := color.New(p.ColorTitle, color.Bold)
	// Title by default does not have a symbol so if user adds one
	// make sure the text is adequately spaced
	if p.SymbolTitle != defaultTitleSymbol {
		text = fmt.Sprintf("\n%s  %s\n", p.SymbolTitle, text)
	} else {
		text = fmt.Sprintf("\n%s\n", text)
	}
	title.Fprint(p.Out, text)
}

func (p *Printer) TitleString(text string) string {
	title := color.New(p.ColorTitle, color.Bold)
	// Title by default does not have a symbol so if user adds one
	// make sure the text is adequately spaced
	if p.SymbolTitle != defaultTitleSymbol {
		text = fmt.Sprintf("%s  %s", p.SymbolTitle, text)
	}
	return title.Sprint(text)
}
