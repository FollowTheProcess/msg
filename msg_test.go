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

// testPrinter returns a default symbols and colors but configured to output to 'out'
// each test should set up their own 'out' from which to read the printed output.
func testPrinter(out io.Writer) *Printer {
	printer := Default()
	printer.Out = out
	return printer
}

// setup returns a testPrinter configured to talk to a bytes.Buffer
// and the pointer to the bytes.Buffer itself to be read from.
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

	got := Default()

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

func TestPrinter_Stitle(t *testing.T) {
	is := is.New(t)
	_, p := setup()

	want := "I'm a Stitle"
	got := p.Stitle("I'm a Stitle")
	is.Equal(got, want)
}

func TestPrinter_StitleSymbol(t *testing.T) {
	is := is.New(t)
	_, p := setup()

	// Change the symbol
	p.SymbolTitle = "💨"

	want := "💨  I'm a Stitle"
	got := p.Stitle("I'm a Stitle")
	is.Equal(got, want)
}

func TestPrinter_Titlef(t *testing.T) {
	is := is.New(t)
	rb, p := setup()

	about := "something"

	want := fmt.Sprintf("\nTitle about: %s\n\n", about)
	p.Titlef("Title about: %s", about)
	is.Equal(rb.String(), want)
}

func TestPrinter_Stitlef(t *testing.T) {
	is := is.New(t)
	_, p := setup()

	about := "something"

	want := fmt.Sprintf("Title about: %s", about)
	got := p.Stitlef("Title about: %s", about)
	is.Equal(got, want)
}

func TestPrinter_Warn(t *testing.T) {
	is := is.New(t)
	rb, p := setup()

	want := fmt.Sprintf("%s  I'm a Warning\n", defaultWarnSymbol)
	p.Warn("I'm a Warning")
	is.Equal(rb.String(), want)
}

func TestPrinter_Warnf(t *testing.T) {
	is := is.New(t)
	rb, p := setup()

	about := "something"

	want := fmt.Sprintf("%s  Warning you about: %s\n", defaultWarnSymbol, about)
	p.Warnf("Warning you about: %s", about)
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

func TestPrinter_Swarn(t *testing.T) {
	is := is.New(t)
	_, p := setup()

	want := fmt.Sprintf("%s  I'm a Swarn", defaultWarnSymbol)
	got := p.Swarn("I'm a Swarn")
	is.Equal(got, want)
}

func TestPrinter_Swarnf(t *testing.T) {
	is := is.New(t)
	_, p := setup()

	about := "something"

	want := fmt.Sprintf("%s  Warning about: %s", defaultWarnSymbol, about)
	got := p.Swarnf("Warning about: %s", about)
	is.Equal(got, want)
}

func TestPrinter_SwarnSymbol(t *testing.T) {
	is := is.New(t)
	_, p := setup()

	// Change the symbol
	p.SymbolWarn = "☢️"

	want := "☢️  I'm a Swarn"
	got := p.Swarn("I'm a Swarn")
	is.Equal(got, want)
}

func TestPrinter_Fail(t *testing.T) {
	is := is.New(t)
	rb, p := setup()

	want := fmt.Sprintf("%s  Error: I'm a Failure\n", defaultFailSymbol)
	p.Fail("I'm a Failure")
	is.Equal(rb.String(), want)
}

func TestPrinter_FailSymbol(t *testing.T) {
	is := is.New(t)
	rb, p := setup()

	// Change the symbol
	p.SymbolFail = "🤬"

	want := "🤬  Error: I'm a Failure\n"
	p.Fail("I'm a Failure")
	is.Equal(rb.String(), want)
}

func TestPrinter_Sfail(t *testing.T) {
	is := is.New(t)
	_, p := setup()

	want := fmt.Sprintf("%s  Error: I'm a Sfail", defaultFailSymbol)
	got := p.Sfail("I'm a Sfail")
	is.Equal(got, want)
}

func TestPrinter_SfailSymbol(t *testing.T) {
	is := is.New(t)
	_, p := setup()

	// Change the symbol
	p.SymbolFail = "🤬"

	want := "🤬  Error: I'm a Sfail"
	got := p.Sfail("I'm a Sfail")
	is.Equal(got, want)
}

func TestPrinter_Sfailf(t *testing.T) {
	is := is.New(t)
	_, p := setup()

	about := "something"

	want := fmt.Sprintf("%s  Error: Something about: %s", defaultFailSymbol, about)
	got := p.Sfailf("Something about: %s", about)
	is.Equal(got, want)
}

func TestPrinter_Failf(t *testing.T) {
	is := is.New(t)
	rb, p := setup()

	about := "something"

	want := fmt.Sprintf("%s  Error: Something: %s\n", defaultFailSymbol, about)
	p.Failf("Something: %s", about)
	is.Equal(rb.String(), want)
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

func TestPrinter_Sgood(t *testing.T) {
	is := is.New(t)
	_, p := setup()

	want := fmt.Sprintf("%s  I'm a Sgood", defaultGoodSymbol)
	got := p.Sgood("I'm a Sgood")
	is.Equal(got, want)
}

func TestPrinter_SgoodSymbol(t *testing.T) {
	is := is.New(t)
	_, p := setup()

	// Change the symbol
	p.SymbolGood = "🎉"

	want := "🎉  I'm a Sgood"
	got := p.Sgood("I'm a Sgood")
	is.Equal(got, want)
}

func TestPrinter_Goodf(t *testing.T) {
	is := is.New(t)
	rb, p := setup()

	about := "something"

	want := fmt.Sprintf("%s  Success: %s\n", defaultGoodSymbol, about)
	p.Goodf("Success: %s", about)
	is.Equal(rb.String(), want)
}

func TestPrinter_Sgoodf(t *testing.T) {
	is := is.New(t)
	_, p := setup()

	about := "something"

	want := fmt.Sprintf("%s  Success: %s", defaultGoodSymbol, about)
	got := p.Sgoodf("Success: %s", about)
	is.Equal(got, want)
}

func TestPrinter_Info(t *testing.T) {
	is := is.New(t)
	rb, p := setup()

	want := fmt.Sprintf("%s  I'm some Info\n", defaultInfoSymbol)
	p.Info("I'm some Info")
	is.Equal(rb.String(), want)
}

func TestPrinter_InfoSymbol(t *testing.T) {
	is := is.New(t)
	rb, p := setup()

	// Change the symbol
	p.SymbolInfo = "🔎"

	want := "🔎  I'm some Info\n"
	p.Info("I'm some Info")
	is.Equal(rb.String(), want)
}

func TestPrinter_Sinfo(t *testing.T) {
	is := is.New(t)
	_, p := setup()

	want := fmt.Sprintf("%s  I'm some Info", defaultInfoSymbol)
	got := p.Sinfo("I'm some Info")
	is.Equal(got, want)
}

func TestPrinter_SinfoSymbol(t *testing.T) {
	is := is.New(t)
	_, p := setup()

	// Change the symbol
	p.SymbolInfo = "🔎"

	want := "🔎  I'm an Sinfo"
	got := p.Sinfo("I'm an Sinfo")
	is.Equal(got, want)
}

func TestPrinter_Infof(t *testing.T) {
	is := is.New(t)
	rb, p := setup()

	about := "something"

	want := fmt.Sprintf("%s  Info: %s\n", defaultInfoSymbol, about)
	p.Infof("Info: %s", about)
	is.Equal(rb.String(), want)
}

func TestPrinter_Sinfof(t *testing.T) {
	is := is.New(t)
	_, p := setup()

	about := "something"

	want := fmt.Sprintf("%s  Info: %s", defaultInfoSymbol, about)
	got := p.Sinfof("Info: %s", about)
	is.Equal(got, want)
}

func TestPrinter_Text(t *testing.T) {
	is := is.New(t)
	rb, p := setup()

	want := fmt.Sprintln("I'm some normal text")
	p.Text("I'm some normal text")
	is.Equal(rb.String(), want)
}

func TestPrinter_Stext(t *testing.T) {
	is := is.New(t)
	_, p := setup()

	want := "I'm some normal text"
	got := p.Stext("I'm some normal text")
	is.Equal(got, want)
}

func TestPrinter_Textf(t *testing.T) {
	is := is.New(t)
	rb, p := setup()

	about := "something"

	want := fmt.Sprintf("Some text about: %s\n", about)
	p.Textf("Some text about: %s", about)
	is.Equal(rb.String(), want)
}

func TestPrinter_Stextf(t *testing.T) {
	is := is.New(t)
	_, p := setup()

	about := "something"

	want := fmt.Sprintf("Text: %s", about)
	got := p.Stextf("Text: %s", about)
	is.Equal(got, want)
}
