---
weight: 20
title: "创建一个应用"
---
# 创建一个应用

每个应用都是通过 langgo.Run 函数开始的

```go
import "github.com/langwan/langgo"

func main() {
    langgo.Run()
}
```

## 调用组件

新建一个 app.yml 文件：

```yaml
hello:
    message: "hello langgo"
```

通过 langgo.Run 函数调用 hello 组件，传入 hello.Instance{} 的具体实例，Langgo 框架会自动加载 app.yml 当中的配置，并绑定 message 给 hello 组件。

通过 hello.Get 函数获取 hello 组件的实例，通过 log.Logger 函数打印 hello 组件的 Message 属性。

```go
import "github.com/langwan/langgo"

func main() {
    langgo.Run(&hello.Instance{})
    log.Logger("app", "main").Info().Str("hello.message", hello.Get().Message).Send()
}
```

获取到的输出：

```shell
hello langgo
```
