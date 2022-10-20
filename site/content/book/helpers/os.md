---
title: "os"
---

# os

系统函数库

## 函数列表

func CopyFile(source string, dest string) (written int64, err error) - 拷贝文件

CreateFolder(p string, ignoreExists bool) error - 创建文件夹，ignoreExists = true 当目标已经存在不会返回错误

FileExists(filename string) bool - 文件是否存在，不区分是否为文件夹

FolderExists(filename string) bool - 文件夹是否存在

GetGoroutineId() uint64 - 获取协程的id号

func MoveFile(source string, dest string) (written int64, err error) - 移动文件

func CopyFileWatcher(source string, dest string, buf []byte, listener IOProgressListener) (written int64, err error) - 拷贝文件并反馈进度

func MoveFileWatcher(source string, dest string, buf []byte, listener IOProgressListener) (written int64, err error) - 移动文件并反馈进度