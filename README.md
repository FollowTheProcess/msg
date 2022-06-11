# msg

[![License](https://img.shields.io/github/license/FollowTheProcess/msg)](https://github.com/FollowTheProcess/msg)
[![Go Report Card](https://goreportcard.com/badge/github.com/FollowTheProcess/msg)](https://goreportcard.com/report/github.com/FollowTheProcess/msg)
[![GitHub](https://img.shields.io/github/v/release/FollowTheProcess/msg?logo=github&sort=semver)](https://github.com/FollowTheProcess/msg)
[![CI](https://github.com/FollowTheProcess/msg/workflows/CI/badge.svg)](https://github.com/FollowTheProcess/msg/actions?query=workflow%3ACI)
[![Go Reference](https://pkg.go.dev/badge/github.com/FollowTheProcess/msg.svg)](https://pkg.go.dev/github.com/FollowTheProcess/msg)

:rocket: A lightweight console printing toolkit for Go CLIs.

* Free software: MIT License

## Project Description

Who else is bored with boring grey text in CLIs? :raising_hand:

We all have fancy terminals, utf-8 is everywhere, no one is still using the stock windows command prompt any more... are they? :raised_eyebrow:

For python CLIs I discovered [ines/wasabi] for this exact purpose and absolutely loved it immediately:

* No dependencies outside the python stdlib
* Configurable if you want but the defaults look great
* Super easy to use and doesn't get in the way of what your CLI is trying to do

Then I started learning go and writing some CLIs I kept wishing I had it here too.

So here it is (mostly) :tada:

I did cheat on the "no dependencies" thing, full credit goes to [fatih/color] for handling all the difficult colouring stuff for me and to [ines/wasabi] for the API design.

![demo](https://github.com/FollowTheProcess/msg/raw/main/img/demo.png)

## Installation

To use `msg` in your code:

```shell
go get github.com/FollowTheProcess/msg@latest
```

## Quickstart

The quickest way to get started is to use the top level functions that come pre-configured with colors and symbols:

``` go
package main

import "github.com/FollowTheProcess/msg"

func main() {
    msg.Title("Your Title here")
    // Do some stuff

    // Give the user an update
    msg.Info("Getting your files")

    // Report success
    msg.Good("It worked!")

    // Uh oh, an error
    msg.Fail("Oh no!, file not found")

    // Warn a user about something
    msg.Warn("This action is irreversible")
}
```

This will get you something that looks like this:

![demo2](https://github.com/FollowTheProcess/msg/raw/main/img/demo2.png)

Not bad! :rocket:

## User Guide

`msg` has 5 types of output:

* **Title** - Section header for a block of output
* **Info** - General info, status updates etc.
* **Good** - To report successful completion of something
* **Warn** - Warn the user about something
* **Fail** - Tell the user there has been a failure of some kind
* **Text** - General purpose normal text

All of them have a single newline appended to the end except `Title` that has a newline before and 2 newlines after, to create separation both from the command you've just run, and from what follows after. You can see how the `Your Title here` is spaced in the example above.

### Print Verbs

`msg` follow Go's lead when it comes to print verbs (think `Printf`, `Sprintf`, `Fprintf` etc.). As such, each of the 5 output types has these methods available (except they have weird names like `Sinfof`, but that's part of the fun.)

For example, let's take the `Warn` message which is the yellow one with the :warning: above:

```go
package main

import "github.com/FollowTheProcess/msg"

func main() {
    
    // The "normal" one -> fmt.Println
    msg.Warn("I'm a Warning")

    // When you want to inject some variables -> fmt.Printf
    // Note: you don't have to put the '\n' at the end, we do that for you
    file := "msg.go"
    msg.Warnf("File: %s not found", file)

    // When you want the string back to use elsewhere -> fmt.Sprint
    // this will strip any newlines and trailing/leading whitespace
    s := msg.Swarn("I'm a Warning string")

    // But I want a string AND formatting -> fmt.Sprintf
    s := msg.Swarnf("Here's a warning string about: %s", file)
}
```

The functions that return strings have all whitespace and newlines trimmed before returning so you'll get back just the raw colored string (and the symbol if you have one, more on this later)

There is no `Fwarnf` though, that's because where `msg` writes to is configured as part of the `Printer`. More on that down here :point_down:

### Configuration

You can get pretty far with just what we've seen up to now, especially if you agree with my choice of symbols and colors.

But I *know* that some of you want finer grained control than this. For you folks, you can instantiate your own `Printer` and have full control over the colors and symbols:

```go
package main

import (
    "os"

    "github.com/FollowTheProcess/msg"
    "github.com/fatih/color"
)

func main() {
    // Defaults shown
    printer := msg.Printer{
        SymbolInfo:  "ℹ",
        SymbolTitle: "",
        SymbolWarn:  "⚠️",
        SymbolFail:  "✘",
        SymbolGood:  "✔",

        // Any fg color from fatih/color will work here
        ColorInfo:  color.FgHiCyan,
        ColorTitle: color.FgCyan,
        ColorWarn:  color.FgYellow,
        ColorFail:  color.FgRed,
        ColorGood:  color.FgGreen,

        // Mainly used for testing but you can set it if you want
        Out: os.Stdout,
    }

    // Now your printer is set up, just call the normal methods
    // all the same methods are available
    printer.Warn("Warning from my custom printer!")
}

```

**But what if I change my mind half way through?**

That's cool, all of the fields in `Printer` are exported so you can just change them whenever you want:

```go
    // Code above omitted

    printer.Info("Old symbol")

    // I want a different symbol for Info for this section only
    printer.SymbolInfo = "🔎"

    printer.Info("New symbol")
```

And you'll get:

![symbol](https://github.com/FollowTheProcess/msg/raw/main/img/symbol.png)

That's about all there is really :tada:

I hope you enjoy it!

## Contributing

### Developing

`msg` is a very simple project and the goal of the project is to remain very simple in line with the good old unix philosophy:

> Write programs that do one thing and do it well.
>
> Ken Thompson

Contributions are very much welcomed but please keep this goal in mind :dart:

`msg` is run as a fairly standard Go project:

* We use all the standard go tools (go test, go build etc.)
* Linting is done with the help of [golangci-lint] (see docs for install help)
* Testing helpers provided by the excellent [matryer/is] package

We use [just] as the command runner (mainly because makefiles make me ill, but also because it's great!)

### Collaborating

No hard and fast rules here but a few guidelines:

* Raise an issue before doing a load of work on a PR, saves everyone bother
* If you add a feature, be sure to add tests to cover what you've added
* If you fix a bug, add a test that would have caught the bug you just squashed
* Be nice :smiley:

### Credits

This package was created with [cookiecutter](https://github.com/cookiecutter/cookiecutter) and the [FollowTheProcess/go_cookie](https://github.com/FollowTheProcess/go_cookie) project template.

[ines/wasabi]: https://github.com/ines/wasabi
[fatih/color]: https://github.com/fatih/color
[golangci-lint]: https://golangci-lint.run
[matryer/is]: https://github.com/matryer/is
[just]: https://github.com/casey/just
