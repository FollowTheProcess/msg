package msg

import (
	"fmt"
	"io"
	"os"
	"strings"

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

// newDefault constructs and returns a default Printer with sensible colors and symbols
func newDefault() *Printer {
	return &Printer{
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

// Title prints a Title message using the configured printer
//
// A Title is distinguishable from all other constructs in msg as it will
// has 1 newline before and 2 newlines after it
//
// If the Printer has a SymbolTitle, it will be prefixed onto 'text'
// with 2 spaces separating them
func (p *Printer) Title(text string) {
	title := color.New(p.ColorTitle, color.Bold)
	// Title by default has an empty string as a symbol
	// sort the spacing out if user sets a symbol
	if p.SymbolTitle != "" {
		text = fmt.Sprintf("\n%s  %s\n\n", p.SymbolTitle, text)
	} else {
		text = fmt.Sprintf("\n%s\n\n", text)
	}
	title.Fprint(p.Out, text)
}

// TitleString is like Title but it returns a string rather than printing it
//
// The returned string will have all it's leading and trailing whitespace/newlines trimmed
// so you have access to the raw string
func (p *Printer) TitleString(text string) string {
	title := color.New(p.ColorTitle, color.Bold)
	// Title by default does not have a symbol so if user adds one
	// make sure the text is adequately spaced
	if p.SymbolTitle != "" {
		text = fmt.Sprintf("%s  %s", p.SymbolTitle, strings.TrimSpace(text))
	}
	return title.Sprint(text)
}

// Warn prints a Warning message using the configured printer
func (p *Printer) Warn(text string) {
	warn := color.New(p.ColorWarn)

	if p.SymbolWarn != "" {
		text = fmt.Sprintf("%s  %s", p.SymbolWarn, text)
	}
	warn.Fprintln(p.Out, text)
}

// WarnString is like Warn but returns a string rather than printing it
func (p *Printer) WarnString(text string) string {
	warn := color.New(p.ColorWarn)
	if p.SymbolWarn != "" {
		text = fmt.Sprintf("%s  %s", p.SymbolWarn, text)
	}
	return warn.Sprint(text)
}

// Title prints a Title message using the default printer
//
// A Title is distinguishable from most other constructs in msg as it will
// always have newlines before and after it
func Title(text string) {
	p := newDefault()
	p.Title(text)
}

// Warn prints a warning message using the default printer
func Warn(text string) {
	p := newDefault()
	p.Warn(text)
}
