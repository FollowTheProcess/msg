// Note we don't actually test for colors in the printed output, this is tested
// extensively in github.com/fatih/color. All we test here are our constructs on top

package msg

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/matryer/is"
)

// testPrinter returns a default printer but configured to output to 'out'
// each test should set up their own 'out' from which to read the printed output
func testPrinter(out io.Writer) Printer {
	printer := Default()
	printer.Out = out
	return printer
}

// setup returns a testPrinter configured to talk to a bytes.Buffer
// and the pointer to the bytes.Buffer itself to be read from
func setup() (*bytes.Buffer, Printer) {
	rb := bytes.NewBuffer(nil)
	p := testPrinter(rb)

	return rb, p
}

func TestDefault(t *testing.T) {
	is := is.New(t)

	want := Printer{
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

	got := Default()

	is.Equal(got, want)
}

func TestTitle(t *testing.T) {
	is := is.New(t)
	rb, p := setup()

	want := "\nI'm a Title\n"
	p.Title("I'm a Title")
	is.Equal(rb.String(), want)
}

func TestTitleSymbol(t *testing.T) {
	is := is.New(t)
	rb, p := setup()

	// Change the symbol
	p.SymbolTitle = "💨"

	want := "\n💨  I'm a Title\n"
	p.Title("I'm a Title")
	is.Equal(rb.String(), want)
}
