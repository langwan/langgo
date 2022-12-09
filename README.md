# Langgo Framework

Langgo 是 go 语言编写的轻量级框架。

## 特性
* 轻量级框架。
* 组件和 helper 的合集
* 支持后端、跨平台应用程序、个人软件开发。
* 容易与其它框架共存。

这里可以查看 

[中文文档](https://langwan.gitbook.io/langgo-v0.5.x/) 

[English Documents](https://langwan.gitbook.io/langgo-v0.5.x/v/english)

![](./logo.png)

安装langgo

```
go get -u github.com/langwan/langgo
```

快速开始

```go
package main

import (
	"github.com/langwan/langgo"
	"github.com/langwan/langgo/components/hello"
	"github.com/langwan/langgo/core/log"
)

func main() {
	langgo.Run(&hello.Instance{Message: "hello"}})
    langgo.Logger("app", "main").Info().Str("hello message", hello.Get().Message).Send()
}
```
## 版本变迁

当前版本为 v0.5.x 文档版本也是 v0.5.x
