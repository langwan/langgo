---
title: "reopen"
---

# reopen

当文件被删除或移动，可以重新创建文件并继续写入，例如 log 收到指令按天切分日志。

## 使用

```go
package main

import (
	"github.com/langwan/langgo/helpers/helper_reopen"
)

func main() {
	p := "./main.log"
	file, err := helper_reopen.NewFileWriter(p)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	file.Write([]byte("langgo"))
	file.Reopen()
	file.Write([]byte("langgo"))
}
```

## 缓冲

NewBufferedFileWriter 会建立一个写入缓存，提高写入能力

```go
package main

import (
	"github.com/langwan/langgo/helpers/helper_reopen"
)

func main() {
	p := "./main.log"
	file, err := helper_reopen.NewFileWriter(p)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	bufferWriter := helper_reopen.NewBufferedFileWriter(file)
	bufferWriter.Write([]byte("langgo"))
	bufferWriter.Reopen()
	bufferWriter.Write([]byte("langgo"))
}
```

默认值

```go
const bufferSize = 256 * 1024
const flushInterval = 30 * time.Second
```

默认缓存大小 256K 默认写入时间 30秒
