---
title: "mysql"
---
# mysql

mysql 内部封装了 grom

## 配置

```
mysql:
  items:
    main:
      conn_max_lifetime: 1h
      dsn: root:123456@tcp(localhost:3306)/simple?charset=utf8mb4&parseTime=True&loc=Local
      max_idle_conns: 1
      max_open_conns: 10
```

允许配置多个数据源。

## 使用

```go
package main

import (
	"fmt"
	"github.com/langwan/langgo"
	"github.com/langwan/langgo/components/mysql"
)

func main() {
	langgo.Run(&mysql.Instance{})
	var one int64
	mysql.Main().Raw("SELECT 1").Scan(&one)
	fmt.Printf("one = %d\n", one)
}
```

## 主数据源和多数据源

调用主数据源

```go
mysql.Main().Raw("SELECT 1").Scan(&one)
```

调用其它数据源使用 mysql.Get 方法 参数为其它数据源配置的 key

```go
mysql.Get("other").Raw("SELECT 1").Scan(&one)
```
