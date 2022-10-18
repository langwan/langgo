---
title: "os"
---

# os

系统函数库

## 函数列表

CopyFile(source string, dest string) error - 拷贝文件

CreateFolder(p string, ignoreExists bool) error - 创建文件夹，ignoreExists = true 当目标已经存在不会返回错误

FileExists(filename string) bool - 文件是否存在，不区分是否为文件夹

FolderExists(filename string) bool - 文件夹是否存在

GetGoroutineId() uint64 - 获取协程的id号

MoveFile(source string, dest string) error - 移动文件
