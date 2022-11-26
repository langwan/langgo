# Langgo Framework

Langgo是一款go语言开发后端和其它应用的多用途框架。

这里可以查看 [全部文档](https://langwan.github.io/langgo)

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
	log.Logger("app", "main").Info().Str("hello message", hello.Get().Message).Send()
}
```
## 版本变迁

当前版本为 v0.5.x 文档版本也是 v0.5.x
