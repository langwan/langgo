---
title: "日志切割"
---

# 日志切割

可以绑定一个系统信号，收到信号后 reopen 日志

```go
package main

import (
	"github.com/langwan/langgo"
	"github.com/langwan/langgo/core/log"
	"syscall"
)

func main() {
	langgo.Run()
	log.SetCuttingSignal(syscall.SIGHUP)
}
```