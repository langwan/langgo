---
title: "设置当前工作目录"
---

## 设置当前工作目录


```go
package main

import (
	"github.com/langwan/langgo"
	"github.com/langwan/langgo/core"
)

func main() {
	core.WorkerDir = "your dir"
	langgo.Run()
}

```

core.WorkerDir - 当前工作目录，默认为启动app的目录。
