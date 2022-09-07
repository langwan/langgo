# Langgo Framework

Langgo是一款go语言开发应用的框架。在B站以视频的形式同步开发。视频地址：https://space.bilibili.com/401571418

## 目录
    - [安装](#安装)

## 安装

基于go 1.19开发

1. 安装langgo
```
go get -u github.com/langwan/langgo
```

2. 导入

```
import "github.com/langwan/langgo"
```

## 快速开始

```
package main

import "github.com/langwan/langgo"

func main() {
    langgo.Init()
    langgo.Run()
}
```