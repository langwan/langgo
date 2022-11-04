---
title: "os"
---

# os

系统函数库

## 函数列表

CreateFolder(p string, ignoreExists bool) error - 创建文件夹，ignoreExists = true 当目标已经存在不会返回错误

FileExists(filename string) bool - 文件是否存在，不区分是否为文件夹

FolderExists(filename string) bool - 文件夹是否存在

GetGoroutineId() uint64 - 获取协程的id号

func CopyFile(dst string, src string) (written int64, err error)  - 拷贝文件

func MoveFile(dst string, src string) (written int64, err error) - 移动文件

func CopyFileWatcher(dst string, src string, buf []byte, listener IOProgressListener) (written int64, err error)  - 拷贝文件并反馈进度

func MoveFileWatcher(dst string, src string, buf []byte, listener IOProgressListener) (written int64, err error) - 移动文件并反馈进度

func GetFileInfo(src string) (fi *FileInfo, err error) - 获取文件的属性 长度 MIME类型 路径 HEAD 

func FileNameWithoutExt(filename string) string - 获取除了扩展名以外的文件名

func TouchFile(p string, ignoreExists bool, createFolder bool) error - 创建空文件