# Langgo Framework

[![Run Tests](https://github.com/langwan/langgo/actions/workflows/go.yml/badge.svg)](https://github.com/langwan/langgo/actions/workflows/go.yml)
![Tag](https://img.shields.io/github/v/tag/langwan/langgo)

## What is Langgo?
Langgo is a go lightweight framework.

## Features
* Lightweight framework.
* Collection of components and helpers.
* Support backend, cross-platform desktop, and personal software development.
* Easy to work with other frameworks.

## Document

[English Document](https://langwan.gitbook.io/langgo-v0.5.x/v/english)

[中文文档](https://langwan.gitbook.io/langgo-v0.5.x/) 

![](./logo.png)

## Examples

[langgo-examples v0.5.x](https://github.com/langwan/langgo-examples/tree/main/0.5.x)

## Installation

```
go get -u github.com/langwan/langgo
```

## Quick Start

```go
package main

import (
	"github.com/langwan/langgo"
	"github.com/langwan/langgo/components/hello"
)

func main() {
	langgo.Run(&hello.Instance{Message: "hello"})
	fmt.Println(hello.Get().Message)
}
```
