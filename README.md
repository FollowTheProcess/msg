# msg

[![License](https://img.shields.io/github/license/FollowTheProcess/msg)](https://github.com/FollowTheProcess/msg)
[![Go Reference](https://pkg.go.dev/badge/github.com/FollowTheProcess/msg.svg)](https://pkg.go.dev/github.com/FollowTheProcess/msg)
[![Go Report Card](https://goreportcard.com/badge/github.com/FollowTheProcess/msg)](https://goreportcard.com/report/github.com/FollowTheProcess/msg)
[![GitHub](https://img.shields.io/github/v/release/FollowTheProcess/msg?logo=github&sort=semver)](https://github.com/FollowTheProcess/msg)
[![CI](https://github.com/FollowTheProcess/msg/workflows/CI/badge.svg)](https://github.com/FollowTheProcess/msg/actions?query=workflow%3ACI)
[![codecov](https://codecov.io/gh/FollowTheProcess/msg/branch/main/graph/badge.svg)](https://codecov.io/gh/FollowTheProcess/msg)

🚀 A lightweight terminal printing toolkit for Go CLIs.

<p align="center">
<img src="https://github.com/FollowTheProcess/msg/raw/main/docs/img/demo.png" alt="demo">
</p>

## Project Description

Who else is bored with boring grey text in CLIs? 🙋🏻‍♂️

We all have fancy terminals, utf-8 is everywhere, no one is still using the stock windows command prompt any more... are they? 🤨

`msg` is a tiny toolkit to make rendering beautiful looking output from CLIs as easy as possible in Go.

It's so easy, you *already* know how it works!

## Installation

```shell
go get github.com/FollowTheProcess/msg@latest
```

## Quickstart

The demo screenshot at the top of the page will get you:

<p align="center">
<img src="https://github.com/FollowTheProcess/msg/raw/main/docs/img/demo.gif" alt="output">
</p>

Not bad! 🚀

## User Guide

`msg` has 5 message types:

* **Title** - Section header for a block of output
* **Info** - General info, status updates etc.
* **Success** - Your CLI has successfully done something
* **Warn** - Warn the user about something e.g. ignoring hidden files
* **Error** - Something has gone wrong in your application

All have a single trailing newline automatically applied except `Title` which has 1 leading and 2 trailing newlines to create separation.

```go
msg.Error("My error message, underlying error: %v", err) // Newlines are handled for you

// i.e. you don't need to do this
msg.Error("My error message, underlying error: %v\n", err)

// Titles are effectively "\n{your title}\n\n"
msg.Title("My title") // Is enough to get separation in sections of output
```

### Stdout/Stderr

By default, every function in `msg` prints to `os.Stdout` with the exception of `msg.Error` which prints to `os.Stderr`.

`msg` also exports "F-style" print functions which can write to any `io.Writer` e.g:

```go
buf := &bytes.Buffer{} // is an io.Writer

msg.Ferror(buf, "My error message")
```

### Credits

This package was created with [copier] and the [FollowTheProcess/go_copier] project template.

[copier]: https://copier.readthedocs.io/en/stable/
[FollowTheProcess/go_copier]: https://github.com/FollowTheProcess/go_copier
