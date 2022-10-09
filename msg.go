// Package msg is a simple, easy to use console printing toolkit for Go CLIs rendering pretty
// formatted output with colours and symbols specific to the particular message type:
//
//   - Info: For general user information and progress updates
//   - Title: Separation between sections of output
//   - Warn: User warnings
//   - Good: Report success
//   - Fail: Report failure (defaults to Stderr)
//
// The symbols and colours used all have sensible defaults but are completely configurable so you
// can tweak the output however you like!
package msg

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/fatih/color"
)

const (
	// Default symbols.
	defaultInfoSymbol  = "ℹ"
	defaultTitleSymbol = ""
	defaultWarnSymbol  = "⚠️"
	defaultFailSymbol  = "✘"
	defaultGoodSymbol  = "✔"

	// Default colors.
	defaultInfoColor  = color.FgHiCyan
	defaultTitleColor = color.FgCyan
	defaultWarnColor  = color.FgYellow
	defaultFailColor  = color.FgHiRed
	defaultGoodColor  = color.FgGreen
)

// Printer is the primary construct in msg, it allows you to configure the colors
// and symbols used for each of the printing methods attached to it.
type Printer struct {
	Stdout      io.Writer       // Stdout
	Stderr      io.Writer       // Stderr
	SymbolInfo  string          // Symbol for the Info output
	SymbolTitle string          // Symbol for the Title output
	SymbolWarn  string          // Symbol for the Warn output
	SymbolFail  string          // Symbol for the Fail output
	SymbolGood  string          // Symbol for the Good output
	ColorInfo   color.Attribute // Color for the Info output
	ColorTitle  color.Attribute // Color for the Title output
	ColorWarn   color.Attribute // Color for the Warn output
	ColorFail   color.Attribute // Color for the Fail output
	ColorGood   color.Attribute // Color for the Good output
}

// Default constructs and returns a default Printer with sensible colors and symbols
// configured to print to os.Stdout.
func Default() *Printer {
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
		Stdout:      os.Stdout,
		Stderr:      os.Stderr,
	}
}

// Title prints a Title message
//
// A Title is distinguishable from all other constructs in msg as it will
// has 1 newline before and 2 newlines after it to create separation
//
// If the Printer has a SymbolTitle, it will be prefixed onto 'text'
// with 2 spaces separating them.
func (p *Printer) Title(text string) {
	title := color.New(p.ColorTitle, color.Bold)
	// Title by default has an empty string as a symbol
	// sort the spacing out if user sets a symbol
	if p.SymbolTitle != "" {
		text = fmt.Sprintf("\n%s  %s\n\n", p.SymbolTitle, text)
	} else {
		text = fmt.Sprintf("\n%s\n\n", text)
	}
	title.Fprint(p.Stdout, text)
}

// Titlef prints a formatted warning message.
func (p *Printer) Titlef(format string, a ...interface{}) {
	text := fmt.Sprintf(format, a...)
	p.Title(text)
}

// Stitle is like Title but it returns a string rather than printing it
//
// The returned string will have all it's leading and trailing whitespace/newlines trimmed
// so you have access to the raw string.
func (p *Printer) Stitle(text string) string {
	title := color.New(p.ColorTitle, color.Bold)
	// Title by default does not have a symbol so if user adds one
	// make sure the text is adequately spaced
	if p.SymbolTitle != "" {
		text = fmt.Sprintf("%s  %s", p.SymbolTitle, strings.TrimSpace(text))
	}
	return title.Sprint(text)
}

// Stitlef returns a formatted title string, stripped of all leading/trailing whitespace.
func (p *Printer) Stitlef(format string, a ...interface{}) string {
	text := fmt.Sprintf(format, a...)
	return p.Stitle(text)
}

// Warn prints a Warning message.
func (p *Printer) Warn(text string) {
	warn := color.New(p.ColorWarn)

	if p.SymbolWarn != "" {
		text = fmt.Sprintf("%s  %s", p.SymbolWarn, text)
	}
	warn.Fprintln(p.Stdout, text)
}

// Warnf prints a formatted warning message.
func (p *Printer) Warnf(format string, a ...interface{}) {
	text := fmt.Sprintf(format, a...)
	p.Warn(text)
}

// Swarn is like Warn but returns a string rather than printing it.
func (p *Printer) Swarn(text string) string {
	warn := color.New(p.ColorWarn)

	if p.SymbolWarn != "" {
		text = fmt.Sprintf("%s  %s", p.SymbolWarn, text)
	}
	return warn.Sprint(text)
}

// Swarnf returns a formatted warning string.
func (p *Printer) Swarnf(format string, a ...interface{}) string {
	text := fmt.Sprintf(format, a...)
	return p.Swarn(text)
}

// Fail prints an error message to Stderr.
func (p *Printer) Fail(text string) {
	failStyle := color.New(p.ColorFail).Add(color.Bold)
	messageStyle := color.New(color.FgHiWhite, color.Bold)

	text = fmt.Sprintf("%s: %s", failStyle.Sprint("Error"), messageStyle.Sprint(text))

	if p.SymbolFail != "" {
		text = fmt.Sprintf("%s  %s", failStyle.Sprint(p.SymbolFail), text)
	}
	fmt.Fprintln(p.Stderr, text)
}

// Failf prints a formatted error message to Stderr.
func (p *Printer) Failf(format string, a ...interface{}) {
	text := fmt.Sprintf(format, a...)
	p.Fail(text)
}

// Sfail is like Fail but returns a string rather than printing it.
func (p *Printer) Sfail(text string) string {
	failStyle := color.New(p.ColorFail).Add(color.Bold)
	messageStyle := color.New(color.FgHiWhite, color.Bold)

	text = fmt.Sprintf("%s: %s", failStyle.Sprint("Error"), messageStyle.Sprint(text))

	if p.SymbolFail != "" {
		text = fmt.Sprintf("%s  %s", failStyle.Sprint(p.SymbolFail), text)
	}
	return text
}

// Sfailf returns a formatted error string.
func (p *Printer) Sfailf(format string, a ...interface{}) string {
	text := fmt.Sprintf(format, a...)
	return p.Sfail(text)
}

// Good prints a success message.
func (p *Printer) Good(text string) {
	good := color.New(p.ColorGood)

	if p.SymbolGood != "" {
		text = fmt.Sprintf("%s  %s", p.SymbolGood, text)
	}
	good.Fprintln(p.Stdout, text)
}

// Goodf prints a formatted success message.
func (p *Printer) Goodf(format string, a ...interface{}) {
	text := fmt.Sprintf(format, a...)
	p.Good(text)
}

// Sgood is like Good but returns a string rather than printing it.
func (p *Printer) Sgood(text string) string {
	good := color.New(p.ColorGood)

	if p.SymbolGood != "" {
		text = fmt.Sprintf("%s  %s", p.SymbolGood, text)
	}
	return good.Sprint(text)
}

// Sgoodf returns a formatted success string.
func (p *Printer) Sgoodf(format string, a ...interface{}) string {
	text := fmt.Sprintf(format, a...)
	return p.Sgood(text)
}

// Info prints an information message.
func (p *Printer) Info(text string) {
	info := color.New(p.ColorInfo)

	if p.SymbolInfo != "" {
		text = fmt.Sprintf("%s  %s", p.SymbolInfo, text)
	}
	info.Fprintln(p.Stdout, text)
}

// Infof prints a formatted information message.
func (p *Printer) Infof(format string, a ...interface{}) {
	text := fmt.Sprintf(format, a...)
	p.Info(text)
}

// Sinfo is like Info but returns a string rather than printing it.
func (p *Printer) Sinfo(text string) string {
	info := color.New(p.ColorInfo)

	if p.SymbolInfo != "" {
		text = fmt.Sprintf("%s  %s", p.SymbolInfo, text)
	}
	return info.Sprint(text)
}

// Sinfof returns a formatted info string.
func (p *Printer) Sinfof(format string, a ...interface{}) string {
	text := fmt.Sprintf(format, a...)
	return p.Sinfo(text)
}

// Text prints a normal, uncoloured message
// you could argue we don't need this as all is does is call fmt.Fprintln but we're here now.
func (p *Printer) Text(text string) {
	fmt.Fprintln(p.Stdout, text)
}

// Textf prints a formatted normal message
// a newline is automatically appended to the end of 'format' so
// you don't have to.
func (p *Printer) Textf(format string, a ...interface{}) {
	fmt.Fprintf(p.Stdout, format+"\n", a...)
}

// Stext is like Text but returns a string rather than printing it.
func (p *Printer) Stext(text string) string {
	return fmt.Sprint(text)
}

// Stextf returns a normal, non coloured formatted string.
func (p *Printer) Stextf(format string, a ...interface{}) string {
	return fmt.Sprintf(format, a...)
}

// Title prints a Title message using the default symbols and colors
//
// A Title is distinguishable from all other constructs in msg as it will
// has 1 newline before and 2 newlines after it.
func Title(text string) {
	p := Default()
	p.Title(text)
}

// Titlef prints a formatted Title message using the default symbols and colors.
func Titlef(format string, a ...interface{}) {
	p := Default()
	p.Titlef(format, a...)
}

// Stitle returns a Title string using the default symbols and colors.
func Stitle(text string) string {
	p := Default()
	return p.Stitle(text)
}

// Stitlef returns a formatted Title string using the default symbols and colors.
func Stitlef(format string, a ...interface{}) string {
	p := Default()
	return p.Stitlef(format, a...)
}

// Warn prints a warning message using the default symbols and colors.
func Warn(text string) {
	p := Default()
	p.Warn(text)
}

// Warnf prints a formatted warning message using the default symbols and colors.
func Warnf(format string, a ...interface{}) {
	p := Default()
	p.Warnf(format, a...)
}

// Swarn returns a warning string using the default symbols and colors.
func Swarn(text string) string {
	p := Default()
	return p.Swarn(text)
}

// Swarnf returns a formatted warning string using the default symbols and colors.
func Swarnf(format string, a ...interface{}) string {
	p := Default()
	return p.Swarnf(format, a...)
}

// Fail prints an error message using the default symbols and colors.
func Fail(text string) {
	p := Default()
	p.Fail(text)
}

// Failf prints a formatted error message using the default symbols and colors.
func Failf(format string, a ...interface{}) {
	p := Default()
	p.Failf(format, a...)
}

// Sfail returns an error message using the default symbols and colors.
func Sfail(text string) string {
	p := Default()
	return p.Sfail(text)
}

// Sfailf returns a formatted error message using the default symbols and colors.
func Sfailf(format string, a ...interface{}) string {
	p := Default()
	return p.Sfailf(format, a...)
}

// Good prints a success message using the default symbols and colors.
func Good(text string) {
	p := Default()
	p.Good(text)
}

// Goodf prints a formatted success message using the default symbols and colors.
func Goodf(format string, a ...interface{}) {
	p := Default()
	p.Goodf(format, a...)
}

// Sgood returns a success message using the default symbols and colors.
func Sgood(text string) string {
	p := Default()
	return p.Sgood(text)
}

// Sgoodf returns a formatted success message using the default symbols and colors.
func Sgoodf(format string, a ...interface{}) string {
	p := Default()
	return p.Sgoodf(format, a...)
}

// Info prints an information message using the default symbols and colors.
func Info(text string) {
	p := Default()
	p.Info(text)
}

// Infof prints a formatted information message using the default symbols and colors.
func Infof(format string, a ...interface{}) {
	p := Default()
	p.Infof(format, a...)
}

// Sinfo returns an information message using the default symbols and colors.
func Sinfo(text string) string {
	p := Default()
	return p.Sinfo(text)
}

// Sinfof returns a formatted information message using the default symbols and colors.
func Sinfof(format string, a ...interface{}) string {
	p := Default()
	return p.Sinfof(format, a...)
}

// Text prints a normal, uncoloured message.
func Text(text string) {
	p := Default()
	p.Text(text)
}

// Textf prints a formatted normal, uncoloured message.
func Textf(format string, a ...interface{}) {
	p := Default()
	p.Textf(format, a...)
}

// Stext returns a normal, uncoloured message.
func Stext(text string) string {
	p := Default()
	return p.Stext(text)
}

// Stextf returns a formatted normal, uncoloured message.
func Stextf(format string, a ...interface{}) string {
	p := Default()
	return p.Stextf(format, a...)
}
