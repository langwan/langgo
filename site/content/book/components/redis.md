---
title: "redis"
---
# redis

redis 内部封装了 go-redis

## 配置

```yaml
redis:
  items:
    main:
      dsn: redis://default:redispw@localhost:55000/0
```

## 使用

```go
package main

import (
	"fmt"
	"github.com/langwan/langgo"
	"github.com/langwan/langgo/components/redis"
	"time"
)

func main() {
	langgo.Run(&redis.Instance{})
	value := "langgo"
	redis.Main().Set("name", value, 10*time.Second)
	get, err := redis.Main().Get("name").Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("value is %s", get)
}
```

## 主数据源和多数据源

调用主数据源

```go
redis.Main().Set("name", value, 10*time.Second)
```

调用其它数据源使用 mysql.Get 方法 参数为其它数据源配置的 key

```go
redis.Get("other").Set("name", value, 10*time.Second)
```
