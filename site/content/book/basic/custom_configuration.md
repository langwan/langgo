---
title: "自定义配置"
---

## 自定义配置

## 结构

```
project_root\
└── main.go
└── app.yml
```

代码

```go
package main

import (
	"github.com/langwan/langgo"
	"github.com/langwan/langgo/core"
	"github.com/langwan/langgo/core/log"
)

type MyConfiguration struct {
	MyName string `yaml:"my_name"`
}

func main() {
	langgo.Run()
	conf := MyConfiguration{}
	core.GetComponentConfiguration("my_configuration", &conf)
	log.Logger("app", "main").Info().Interface("conf", conf).Send()
}
```
core.GetComponentConfiguration 方法可以把自定义配置绑定到对象上。

