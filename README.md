# msg

[![License](https://img.shields.io/github/license/FollowTheProcess/msg)](https://github.com/FollowTheProcess/msg)
[![Go Report Card](https://goreportcard.com/badge/github.com/FollowTheProcess/msg)](https://goreportcard.com/report/github.com/FollowTheProcess/msg)
[![GitHub](https://img.shields.io/github/v/release/FollowTheProcess/msg?logo=github&sort=semver)](https://github.com/FollowTheProcess/msg)
[![CI](https://github.com/FollowTheProcess/msg/workflows/CI/badge.svg)](https://github.com/FollowTheProcess/msg/actions?query=workflow%3ACI)

:rocket: A lightweight console printing toolkit for Go CLIs.

* Free software: MIT License

## Project Description

When writing [pytoil] I discovered [ines/wasabi] and absolutely loved it immediately! Then when I started learning go and writing some CLI's I kept wishing I had it there too.

So here it is :tada:

Full credit goes to [fatih/color] for handling all the difficult colouring stuff for me and to [ines/wasabi] for the API design.

![demo](https://github.com/FollowTheProcess/msg/raw/main/img/demo.png)

## Installation

To use `msg` in your code:

```shell
go get github.com/FollowTheProcess/msg@latest
```

## Quickstart

The quickest way to get started is to use the pre-configured functions:

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

### Credits

This package was created with [cookiecutter](https://github.com/cookiecutter/cookiecutter) and the [FollowTheProcess/go_cookie](https://github.com/FollowTheProcess/go_cookie) project template.

[pytoil]: https://github.com/FollowTheProcess/pytoil
[ines/wasabi]: https://github.com/ines/wasabi
[fatih/color]: https://github.com/fatih/color
