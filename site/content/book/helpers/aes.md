---
title: "aes"
---

# aes

## 使用

加密解密

```go
encrypt, err := helper_aes.Encrypt(key, data)
decrypt, err := helper_aes.Decrypt(key, encrypt)
```

## 函数列表

Encrypt(key, src []byte) (data []byte, err error) - 解密数据

Decrypt(key, src []byte) (data []byte, err error) - 加密数据
