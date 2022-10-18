---
title: "jwt"
---
# jwt

实现了JWT认证

## 配置

```yaml
jwt:
    secret: "123456"
```

## 使用

```go
func main() {
	langgo.Run(&jwt.Instance{})
	payload := struct {
		Name string
	}{Name: "langwan"}
	sign, err := jwt.Sign(payload)
	if err != nil {
		panic(err)
	}
	err = jwt.Verify(sign)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("verify ok")
	}
}
```

## 方法

`Sign` - 签名 成功返回签名 失败返回错误。

`Verify` - 验证签名 成功 err = nil，失败返回错误。
