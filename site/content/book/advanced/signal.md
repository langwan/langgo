---
title: "信号处理"
---

# 信号处理

对 go 原生的信号处理做了更好的封装

绑定信号处理程序

```go
package main

import (
	"fmt"
	"github.com/langwan/langgo"
	"github.com/langwan/langgo/core"
	"os"
	"syscall"
)

func main() {
	langgo.Run()
	done := make(chan struct{})
	core.SignalHandlers(func(sig os.Signal) {
		fmt.Printf("sig = %d\n", sig)
	}, syscall.SIGHUP, syscall.SIGUSR1)
	core.SignalNotify()
	fmt.Printf("pid = %d\n", os.Getpid())
	<-done
}
```

SignalHandlers(handler func(sig os.Signal), signals ...os.Signal) - 绑定一个或多个信号到同一个函数
SignalNotify() - 异步等待信号并处理，该函数只有第一次调用有效，无需多次调用


输入指令

```
kill -HUP 25116
kill -USR1 25116
```

输出

```
pid = 25116
sig = 1
sig = 30
```