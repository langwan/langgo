---
title: "快速上手"
---

# 快速上手

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

只需要使用 langgo.Run 函数就可以初始化 langgo 框架

hello 是自带的最简单的组件
