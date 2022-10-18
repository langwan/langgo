---
weight: 22
title: "日志"
---

# 日志

log 内部封装了 [zerolog](https://github.com/rs/zerolog)

## 使用

```go
import "github.com/langwan/langgo"

func main() {
    langgo.Run(&hello.Instance{})
    log.Logger("app", "main").Info().Str("hello.message", hello.Get().Message).Send()
}
```

app 是日志的名称 main 是tag，Info() 是当前日志的级别，`log.Logger("app", "main")`代码之后的部分完全可以参考 [zerolog](https://github.com/rs/zerolog) 使用方法。

日志会自动被创建在项目中 logs 子目录当中。

```
project_root\
└── logs\
    └── app.log
└── main.go
└── app.yml
```

