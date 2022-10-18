---
title: "gen"
---

# gen

gen 生成库，例如 uuid 随机字符串等。

## 函数列表

Uuid() string - 产生标准 uuid

UuidShort() string - 产生短格式的 uuid

UuidNoSeparator() string - 生成没有分隔符的 uuid

RandInt(min, max int64) int64 - 随机生成 int

RandString(n int, letters ...string) (string, error) - 随机生成字符串
