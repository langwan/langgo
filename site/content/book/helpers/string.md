---
title: "string"
---

# string

字符串函数库

## 函数列表

IsEmpty(val string) bool - 去除左右空白之后判断字符串是否为空，是返回true
func Utf8StringLength(str string) int - 获取utf8字符串的长度
func Utf8TruncateText(text string, max int, omission string) string - 从utf8字符串中截取内容