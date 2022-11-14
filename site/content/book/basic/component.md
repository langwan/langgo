---
weight: 21
title: "组件"
---
# 组件

组件的特征是必须创建了一个实例，并常驻内存当中。

## 组件的属性

代码结构

绑定组件的属性有两种方式


```
project_root\
└── main.go
└── app.yml
```

方式一 通过配置文件加载 app.yml

```yaml
hello:
    message: "hello"
```

方式二 实例化的时候传入

```go
langgo.Run(&hello.Instance{Message: "hello"}})
```

方式三 框架初始化以后加载

```go
core.LoadComponents(&hello.Instance{})
```

## 内置组件

这些组件是 langgo 框架内置的组件，例如 mysql、redis、sqlite 使用这些组件可以搭建应用系统。

## 自定义组件

结构

```
project_root\
└── components
    └── my\
        └── my.go
└── main.go
└── app.yml
```

在自己项目的实现 core.Component 接口

```go
package my

type Instance struct {
	Message string `yaml:"message"`
}

const name = "my"

var instance *Instance

func (i *Instance) Run() error {
	instance = i
	return nil
}

func (i *Instance) GetName() string {
	return name
}

func Get() *Instance {
	return instance
}

```

配置

```yaml
my:
  message: "im custom component"
```

使用

```go
package main

import (
	"fmt"
	"github.com/langwan/langgo"
	"github.com/langwan/langgo/examples/component/custom/my"
)

func main() {
	langgo.Run(&my.Instance{})
	fmt.Printf("component name is `%s`, message is `%s`\n", my.Get().GetName(), my.Get().Message)
}
```
