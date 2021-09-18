package main

import (
	"fmt"

	"github.com/FollowTheProcess/msg/msg"
)

func main() {
	printer := msg.Default()

	printer.Title("I'm a title")
	fmt.Println("I'm some text on the next line")

	s := printer.TitleString("I'm a titlestring")
	fmt.Println(s)

	printer.SymbolTitle = "💨"

	s = printer.TitleString("I'm a titlestring with a symbol")
	fmt.Println(s)
}
