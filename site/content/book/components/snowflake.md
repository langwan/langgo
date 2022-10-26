---
title: "snowflake"
---
# snowflake

实现了snowflake id

## 配置

```yaml
snowflake:
  machine_id: 1
```

机器编号范围 0 - 1023

## 使用

```go
func main() {
    langgo.Run(&snowflake.Instance{})
	id := snowflake.Gen()
    fmt.Printf("Int64  ID: %d\n", id)
}
```

## 方法

`Gen` - 返回snowflake id
