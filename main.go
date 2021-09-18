package main

import (
	"fmt"

	"github.com/FollowTheProcess/msg/hello"
)

var Version = "dev"

func main() {
	message := hello.Hello()

	fmt.Println(message)
	fmt.Println("Version:", Version)
}
