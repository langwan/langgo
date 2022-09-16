# Langgo Framework

Langgo是一款go语言开发应用的框架。在B站以视频的形式同步开发。视频地址：https://space.bilibili.com/401571418/channel/collectiondetail?sid=699075

## 目录

 - [安装](#安装)
 - [快速开始](#快速开始)
 - [开发视频](#开发视频)
 - [grpc](#grpc)
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

import (
	"github.com/langwan/langgo"
	"github.com/langwan/langgo/components/hello"
	"github.com/langwan/langgo/core/log"
)

func main() {
	langgo.Run(&hello.Instance{Message: "hello component"})
	log.Logger("component", "hello").Info().Msg(hello.GetInstance().Message)
}
```

## 开发视频

视频地址 https://space.bilibili.com/401571418/channel/collectiondetail?sid=699075

## grpc

grpc支持单机模式和etcd服务发现两种模式，可以参考examples/grpc/single和examples/grpc/etcd两个例子。