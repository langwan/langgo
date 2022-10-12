# Langgo Framework

Langgo是一款go语言开发应用的框架。在B站以视频的形式同步开发。视频地址：https://space.bilibili.com/401571418/channel/collectiondetail?sid=699075

## 目录

 - [安装](#安装)
 - [快速开始](#快速开始)
 - [开发视频](#开发视频)
 - [核心组件]
   - [日志](#日志)
 - [组件](#组件)
   - [mysql](#mysql)
   - [redis](#redis)
 - [helper](#helper) 
   - [rsa](#rsa)
   - [aes](#aes)
   - [grpc](#grpc)
 - [自定义组件](#自定义组件)
 - [性能测试]
   - [日志写入速度](#日志写入速度)
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

```go
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

## 组件

组件有两个特征，需要配置或者需要持久化的实例，例如`mysql`，需要对数据库连接参数进行配置，连接也需要持久化在内存当中，这两种情况下需要制作成组件。`langgo`有一些内置的组件。


## 开发视频

视频地址 https://space.bilibili.com/401571418/channel/collectiondetail?sid=699075

## 日志

日志扩展自 [zerolog](https://github.com/rs/zerolog)

特性：

支持自定义信号切割日志

在配置文件中增加log配置
```yaml
log:
  reopen_signal: 31
```
`reopen_signal`是自定义信号量的int值，当服务收到信号以后会自动mv出新的日志文件，方便按照日、小时等自定义策略来切分日志，mv不会丢失日志。

2. 支持动态创建日志文件例如 app.log order.log 方便日志分类

```go
log.Logger("app", "login").Info().Interface("request", request).Send()
log.Logger("order", "create").Info().Interface("request", request).Send()
```
`app`, `order`会自动被创建为`logs/app.log`, `logs/order.log`



## grpc

grpc支持单机模式和etcd服务发现两种模式，可以参考examples/grpc/single和examples/grpc/etcd两个例子。

完成特性：

* 支持etcd等服务发现模式
* 支持tls双向认证
* 支持链表式多中间件

## mysql

参考 `examples/mysql`，实际整合了`gorm`

mysql配置支持多个mysql账号，例如：

```yaml
mysql:
  main:
    dsn: main:123456@tcp(localhost:3306)/simple?charset=utf8mb4&parseTime=True&loc=Local
    conn_max_lifetime: 1h
    max_idle_conns: 1
    max_open_conns: 10
  order:
    dsn: order:123456@tcp(localhost:3306)/simple?charset=utf8mb4&parseTime=True&loc=Local
    conn_max_lifetime: 1h
    max_idle_conns: 1
    max_open_conns: 10

```

这样可以支持项目会拥有多个mysql数据库

```go
langgo.Run(&mysql.Instance{})
var one int
mysql.Main().Raw("SELECT 1").Scan(&one)
fmt.Println(one)
```

`mysql.Main()` 表示获取配置中`main`下的mysql配置，如果想获取`order`，需要使用 `mysql.Get("order")`

## redis

与mysql差不多，可以对redis进行配置并使用多个redis数据源，可以对redis连接池进行配置，实际上整合了`go-reids`

## 自定义组件

完全可以在自己的项目中实现自定义组件与使用`langgo`框架中的内置组件的方式一样，只需要您在自己的项目里拷贝出`hello`组件的代码，进行改造，然后在启动的时候放在`langgo.Run`里一起启动就可以了：

```go
langgo.Run(&my.Instance{})
```

组件的配置文件放在`app.yml`配置中就可以了，例如：
```yaml
my:
  message: my name is langgo
```

可以参考 examples/component/custom 这个例子

## helper

无需配置和内存持久化，直接调用的函数集合。

## rsa

支持 加密、解密、创建密钥（根据长度）、签名、校验等方法

## aes 

支持加密、解密等方法

## 日志写入速度

在`core/log/log_test.go`中有相关的基准测试方法，我们跑的结果如下：

```
BenchmarkLoggerSystemLog
BenchmarkLoggerSystemLog-8     	  540446	      2113 ns/op
BenchmarkZerologFile
BenchmarkZerologFile-8         	  508726	      2353 ns/op
BenchmarkZerologConsole
BenchmarkZerologConsole-8      	  228208	      5316 ns/op
BenchmarkLoggerLanggo
BenchmarkLoggerLanggo-8        	  523492	      2409 ns/op
BenchmarkLoggerLanggoMulti
BenchmarkLoggerLanggoMulti-8   	  446125	      2660 ns/op
```
* BenchmarkLoggerSystemLog-8 使用go自带的log写入文件 大概是54万次/秒
* BenchmarkZerologFile-8 使用zerolog直接写入文件 大概是 50万次/秒
* BenchmarkZerologConsole-8 使用zerolog的consoleWriter包装以后写入文件 大概是 22万次/秒 关于zerolog的这个问题，请参考我的B站视频 [优化日志写入速度](https://www.bilibili.com/video/BV1XN4y1c7Jj/?spm_id_from=333.999.0.0)
* BenchmarkLoggerLanggo-8 使用langgo的日志（二次包装zerolog直接写入文件） 大概是 52万次/秒
* BenchmarkLoggerLanggoMulti-8 使用langgo的动态生成日志文件功能 大概是 44万次/秒