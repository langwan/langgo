---
title: "size"
---

# size

从人类可读字符串当中解析字节数，或者字节数格式化成人类可读的字符串

## 函数列表

BytesSize(size float64) string - 把数字转换成人类可读的字符串，基础单位是二进制例如1024=1k

CustomSize(format string, size float64, base float64, _map []string) string - 格式化数字

RAMInBytes(size string) (int64, error) - 把人类可读的字符串转换成int64，基础单位是二进制，例如1k=1024

FromHumanSize(size string) (int64, error) - 把人类可读的字符串转换车鞥int64，基础单位是10进制，例如1k=1000
