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
	defaultInfoColor  = color.FgHiCyan
	defaultTitleColor = color.FgCyan
	defaultWarnColor  = color.FgYellow
	defaultFailColor  = color.FgRed
	defaultGoodColor  = color.FgGreen
)

// Printer is the primary construct in msg, it allows you to configure the colors
// and symbols used for each of the printing methods attached to it.
type Printer struct {
	// Symbols

	// The symbol used to prefix Info and Sinfo
	SymbolInfo string
	// The symbol used to prefix Title and Stitle
	SymbolTitle string
	// The symbol used ot prefix Warn and Swarn
	SymbolWarn string
	// The symbol used to prefix Fail and Sfail
	SymbolFail string
	// The symbol used to prefix Good and Sgood
	SymbolGood string

	// Colors

	// The color used for Info and Sinfo
	ColorInfo color.Attribute
	// The color used for Title and Stitle
	ColorTitle color.Attribute
	// The color used for Warn and Swarn
	ColorWarn color.Attribute
	// The color used for Fail and Sfail
	ColorFail color.Attribute
	// The color used for Good and Sgood
	ColorGood color.Attribute

	// Output

	// Where to do the printing, useful for testing
	Out io.Writer
}

// newDefault constructs and returns a default symbols and colors with sensible colors and symbols
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

// Title prints a Title message
//
// A Title is distinguishable from all other constructs in msg as it will
// has 1 newline before and 2 newlines after it to create separation
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

// Titlef prints a formatted warning message
func (p *Printer) Titlef(format string, a ...interface{}) {
	text := fmt.Sprintf(format, a...)
	p.Title(text)
}

// Stitle is like Title but it returns a string rather than printing it
//
// The returned string will have all it's leading and trailing whitespace/newlines trimmed
// so you have access to the raw string
func (p *Printer) Stitle(text string) string {
	title := color.New(p.ColorTitle, color.Bold)
	// Title by default does not have a symbol so if user adds one
	// make sure the text is adequately spaced
	if p.SymbolTitle != "" {
		text = fmt.Sprintf("%s  %s", p.SymbolTitle, strings.TrimSpace(text))
	}
	return title.Sprint(text)
}

// Stitlef returns a formatted title string, stripped of all leading/trailing whitespace
func (p *Printer) Stitlef(format string, a ...interface{}) string {
	text := fmt.Sprintf(format, a...)
	return p.Stitle(text)
}

// Warn prints a Warning message
func (p *Printer) Warn(text string) {
	warn := color.New(p.ColorWarn)

	if p.SymbolWarn != "" {
		text = fmt.Sprintf("%s  %s", p.SymbolWarn, text)
	}
	warn.Fprintln(p.Out, text)
}

// Warnf prints a formatted warning message
func (p *Printer) Warnf(format string, a ...interface{}) {
	text := fmt.Sprintf(format, a...)
	p.Warn(text)
}

// Swarn is like Warn but returns a string rather than printing it
func (p *Printer) Swarn(text string) string {
	warn := color.New(p.ColorWarn)

	if p.SymbolWarn != "" {
		text = fmt.Sprintf("%s  %s", p.SymbolWarn, text)
	}
	return warn.Sprint(text)
}

// Swarnf returns a formatted warning string
func (p *Printer) Swarnf(format string, a ...interface{}) string {
	text := fmt.Sprintf(format, a...)
	return p.Swarn(text)
}

// Fail prints an error message
func (p *Printer) Fail(text string) {
	fail := color.New(p.ColorFail)

	if p.SymbolWarn != "" {
		text = fmt.Sprintf("%s  %s", p.SymbolFail, text)
	}
	fail.Fprintln(p.Out, text)
}

// Failf prints a formatted error message
func (p *Printer) Failf(format string, a ...interface{}) {
	text := fmt.Sprintf(format, a...)
	p.Fail(text)
}

// Sfail is like Fail but returns a string rather than printing it
func (p *Printer) Sfail(text string) string {
	fail := color.New(p.ColorFail)

	if p.SymbolWarn != "" {
		text = fmt.Sprintf("%s  %s", p.SymbolFail, text)
	}
	return fail.Sprint(text)
}

// Sfailf returns a formatted error string
func (p *Printer) Sfailf(format string, a ...interface{}) string {
	text := fmt.Sprintf(format, a...)
	return p.Sfail(text)
}

// Good prints a success message
func (p *Printer) Good(text string) {
	good := color.New(p.ColorGood)

	if p.SymbolGood != "" {
		text = fmt.Sprintf("%s  %s", p.SymbolGood, text)
	}
	good.Fprintln(p.Out, text)
}

// Goodf prints a formatted success message
func (p *Printer) Goodf(format string, a ...interface{}) {
	text := fmt.Sprintf(format, a...)
	p.Good(text)
}

// Sgood is like Good but returns a string rather than printing it
func (p *Printer) Sgood(text string) string {
	good := color.New(p.ColorGood)

	if p.SymbolGood != "" {
		text = fmt.Sprintf("%s  %s", p.SymbolGood, text)
	}
	return good.Sprint(text)
}

// Sgoodf returns a formatted success string
func (p *Printer) Sgoodf(format string, a ...interface{}) string {
	text := fmt.Sprintf(format, a...)
	return p.Sgood(text)
}

// Info prints an information message
func (p *Printer) Info(text string) {
	info := color.New(p.ColorInfo)

	if p.SymbolInfo != "" {
		text = fmt.Sprintf("%s  %s", p.SymbolInfo, text)
	}
	info.Fprintln(p.Out, text)
}

// Infof prints a formatted information message
func (p *Printer) Infof(format string, a ...interface{}) {
	text := fmt.Sprintf(format, a...)
	p.Info(text)
}

// Sinfo is like Info but returns a string rather than printing it
func (p *Printer) Sinfo(text string) string {
	info := color.New(p.ColorInfo)

	if p.SymbolInfo != "" {
		text = fmt.Sprintf("%s  %s", p.SymbolInfo, text)
	}
	return info.Sprint(text)
}

// Sinfof returns a formatted info string
func (p *Printer) Sinfof(format string, a ...interface{}) string {
	text := fmt.Sprintf(format, a...)
	return p.Sinfo(text)
}

// Text prints a normal, uncoloured message
// you could argue we don't need this as all is does is call fmt.Fprintln but we're here now
func (p *Printer) Text(text string) {
	fmt.Fprintln(p.Out, text)
}

// Textf prints a formatted normal message
// a newline is automatically appended to the end of 'format' so
// you don't have to
func (p *Printer) Textf(format string, a ...interface{}) {
	fmt.Fprintf(p.Out, format+"\n", a...)
}

// Stext is like Text but returns a string rather than printing it
func (p *Printer) Stext(text string) string {
	return fmt.Sprint(text)
}

// Stextf returns a normal, non coloured formatted string
func (p *Printer) Stextf(format string, a ...interface{}) string {
	return fmt.Sprintf(format, a...)
}

// Title prints a Title message using the default symbols and colors
//
// A Title is distinguishable from all other constructs in msg as it will
// has 1 newline before and 2 newlines after it
func Title(text string) {
	p := newDefault()
	p.Title(text)
}

// Titlef prints a formatted Title message using the default symbols and colors
func Titlef(format string, a ...interface{}) {
	p := newDefault()
	p.Titlef(format, a...)
}

// Stitle returns a Title string using the default symbols and colors
func Stitle(text string) string {
	p := newDefault()
	return p.Stitle(text)
}

// Stitlef returns a formatted Title string using the default symbols and colors
func Stitlef(format string, a ...interface{}) string {
	p := newDefault()
	return p.Stitlef(format, a...)
}

// Warn prints a warning message using the default symbols and colors
func Warn(text string) {
	p := newDefault()
	p.Warn(text)
}

// Warnf prints a formatted warning message using the default symbols and colors
func Warnf(format string, a ...interface{}) {
	p := newDefault()
	p.Warnf(format, a...)
}

// Swarn returns a warning string using the default symbols and colors
func Swarn(text string) string {
	p := newDefault()
	return p.Swarn(text)
}

// Swarnf returns a formatted warning string using the default symbols and colors
func Swarnf(format string, a ...interface{}) string {
	p := newDefault()
	return p.Swarnf(format, a...)
}

// Fail prints an error message using the default symbols and colors
func Fail(text string) {
	p := newDefault()
	p.Fail(text)
}

// Failf prints a formatted error message using the default symbols and colors
func Failf(format string, a ...interface{}) {
	p := newDefault()
	p.Failf(format, a...)
}

// Sfail returns an error message using the default symbols and colors
func Sfail(text string) string {
	p := newDefault()
	return p.Sfail(text)
}

// Sfailf returns a formatted error message using the default symbols and colors
func Sfailf(format string, a ...interface{}) string {
	p := newDefault()
	return p.Sfailf(format, a...)
}

// Good prints a success message using the default symbols and colors
func Good(text string) {
	p := newDefault()
	p.Good(text)
}

// Goodf prints a formatted success message using the default symbols and colors
func Goodf(format string, a ...interface{}) {
	p := newDefault()
	p.Goodf(format, a...)
}

// Sgood returns a success message using the default symbols and colors
func Sgood(text string) string {
	p := newDefault()
	return p.Sgood(text)
}

// Sgoodf returns a formatted success message using the default symbols and colors
func Sgoodf(format string, a ...interface{}) string {
	p := newDefault()
	return p.Sgoodf(format, a...)
}

// Info prints an information message using the default symbols and colors
func Info(text string) {
	p := newDefault()
	p.Info(text)
}

// Infof prints a formatted information message using the default symbols and colors
func Infof(format string, a ...interface{}) {
	p := newDefault()
	p.Infof(format, a...)
}

// Sinfo returns an information message using the default symbols and colors
func Sinfo(text string) string {
	p := newDefault()
	return p.Sinfo(text)
}

// Sinfof returns a formatted information message using the default symbols and colors
func Sinfof(format string, a ...interface{}) string {
	p := newDefault()
	return p.Sinfof(format, a...)
}

// Text prints a normal, uncoloured message
func Text(text string) {
	p := newDefault()
	p.Text(text)
}

// Textf prints a formatted normal, uncoloured message
func Textf(format string, a ...interface{}) {
	p := newDefault()
	p.Textf(format, a...)
}

// Stext returns a normal, uncoloured message
func Stext(text string) string {
	p := newDefault()
	return p.Stext(text)
}

// Stextf returns a formatted normal, uncoloured message
func Stextf(format string, a ...interface{}) string {
	p := newDefault()
	return p.Stextf(format, a...)
}
