---
title: "开发和正式环境"
---

## 开发环境

可以通过设置 env 的方式在 正式和开发环境中切换，默认为 development
```
langgo_env=production ./app
langgo_env=development ./app
```

代码

```go
package main

import (
	"fmt"
	"github.com/langwan/langgo"
	"github.com/langwan/langgo/core"
)

func main() {
	langgo.Run()
	fmt.Printf("mode = %s\n", core.EnvName)
}
```

langgo_env=production ./app 输出

```
mode = production
```


./app 输出

```
mode = development
```