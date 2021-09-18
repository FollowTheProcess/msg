// Note we don't actually test for colors in the printed output, this is tested
// extensively in github.com/fatih/color. All we test here are our constructs on top

package msg

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/matryer/is"
)

// testPrinter returns a default printer but configured to output to 'out'
// each test should set up their own 'out' from which to read the printed output
func testPrinter(out io.Writer) *Printer {
	printer := newDefault()
	printer.Out = out
	return printer
}

// setup returns a testPrinter configured to talk to a bytes.Buffer
// and the pointer to the bytes.Buffer itself to be read from
func setup() (*bytes.Buffer, *Printer) {
	rb := bytes.NewBuffer(nil)
	p := testPrinter(rb)

	return rb, p
}

func TestNewDefault(t *testing.T) {
	is := is.New(t)

	want := &Printer{
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

	got := newDefault()

	is.Equal(got, want)
}

func TestPrinter_Title(t *testing.T) {
	is := is.New(t)
	rb, p := setup()

	want := "\nI'm a Title\n\n"
	p.Title("I'm a Title")
	is.Equal(rb.String(), want)
}

func TestPrinter_TitleSymbol(t *testing.T) {
	is := is.New(t)
	rb, p := setup()

	// Add a symbol
	p.SymbolTitle = "💨"

	want := "\n💨  I'm a Title\n\n"
	p.Title("I'm a Title")
	is.Equal(rb.String(), want)
}

func TestPrinter_TitleString(t *testing.T) {
	is := is.New(t)
	_, p := setup()

	want := "I'm a Titlestring"
	got := p.TitleString("I'm a Titlestring")
	is.Equal(got, want)
}

func TestPrinter_TitleStringSymbol(t *testing.T) {
	is := is.New(t)
	_, p := setup()

	// Change the symbol
	p.SymbolTitle = "💨"

	want := "💨  I'm a Titlestring"
	got := p.TitleString("I'm a Titlestring")
	is.Equal(got, want)
}

func TestPrinter_Warn(t *testing.T) {
	is := is.New(t)
	rb, p := setup()

	want := fmt.Sprintf("%s  I'm a Warning\n", defaultWarnSymbol)
	p.Warn("I'm a Warning")
	is.Equal(rb.String(), want)
}

func TestPrinter_WarnSymbol(t *testing.T) {
	is := is.New(t)
	rb, p := setup()

	// Change the symbol
	p.SymbolWarn = "☢️"

	want := "☢️  I'm a Warning\n"
	p.Warn("I'm a Warning")
	is.Equal(rb.String(), want)
}

func TestPrinter_WarnString(t *testing.T) {
	is := is.New(t)
	_, p := setup()

	want := fmt.Sprintf("%s  I'm a Warnstring", defaultWarnSymbol)
	got := p.WarnString("I'm a Warnstring")
	is.Equal(got, want)
}

func TestPrinter_WarnStringSymbol(t *testing.T) {
	is := is.New(t)
	_, p := setup()

	// Change the symbol
	p.SymbolWarn = "☢️"

	want := "☢️  I'm a Warnstring"
	got := p.WarnString("I'm a Warnstring")
	is.Equal(got, want)
}

func TestPrinter_Fail(t *testing.T) {
	is := is.New(t)
	rb, p := setup()

	want := fmt.Sprintf("%s  I'm a Failure\n", defaultFailSymbol)
	p.Fail("I'm a Failure")
	is.Equal(rb.String(), want)
}

func TestPrinter_FailSymbol(t *testing.T) {
	is := is.New(t)
	rb, p := setup()

	// Change the symbol
	p.SymbolFail = "🤬"

	want := "🤬  I'm a Failure\n"
	p.Fail("I'm a Failure")
	is.Equal(rb.String(), want)
}

func TestPrinter_FailString(t *testing.T) {
	is := is.New(t)
	_, p := setup()

	want := fmt.Sprintf("%s  I'm a Failstring", defaultFailSymbol)
	got := p.FailString("I'm a Failstring")
	is.Equal(got, want)
}

func TestPrinter_FailStringSymbol(t *testing.T) {
	is := is.New(t)
	_, p := setup()

	// Change the symbol
	p.SymbolFail = "🤬"

	want := "🤬  I'm a Failstring"
	got := p.FailString("I'm a Failstring")
	is.Equal(got, want)
}

func TestPrinter_Good(t *testing.T) {
	is := is.New(t)
	rb, p := setup()

	want := fmt.Sprintf("%s  I'm a Success\n", defaultGoodSymbol)
	p.Good("I'm a Success")
	is.Equal(rb.String(), want)
}

func TestPrinter_GoodSymbol(t *testing.T) {
	is := is.New(t)
	rb, p := setup()

	// Change the symbol
	p.SymbolGood = "🎉"

	want := "🎉  I'm a Success\n"
	p.Good("I'm a Success")
	is.Equal(rb.String(), want)
}

func TestPrinter_GoodString(t *testing.T) {
	is := is.New(t)
	_, p := setup()

	want := fmt.Sprintf("%s  I'm a Goodstring", defaultGoodSymbol)
	got := p.GoodString("I'm a Goodstring")
	is.Equal(got, want)
}

func TestPrinter_GoodStringSymbol(t *testing.T) {
	is := is.New(t)
	_, p := setup()

	// Change the symbol
	p.SymbolGood = "🎉"

	want := "🎉  I'm a Goodstring"
	got := p.GoodString("I'm a Goodstring")
	is.Equal(got, want)
}
