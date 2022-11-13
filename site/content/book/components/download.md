---
title: "download"
---
# download

多协程分片文件下载

例如下载一个11MB的文件，partSize=2MB，会被切分成6个分片part，放在协程中进行下载，协程的数量等于workers，例如workers=5表示，6个part在5个协程中排队下载。下载过程中会产生xxx.db、xxx.dp0、xxx.dp1的临时文件，当最后一个分片被下载完成，会统一合并成一个文件，并删除这些中间的文件。

为了减少重复代码的开发，协程池采用 [ants](github.com/panjf2000/ants) 组件实现

## 配置

```yaml
download:
  workers: 5
  partSize: 5m
  bufSize: 200k
```

workers - 协程数
partSize - 分片大小
bufSize - 写文件时候的缓存大小

## 使用

```go
func main() {
	url := "https://xxx/xxx.mp4"
    langgo.Run(&download.Instance{})
    httpReader := HttpReader{Url: url}
    download.Get().Download(context.Background(), url, "./example.mp4", &httpReader, &Listener{})
}
```

## reader

为了支持更多的协议，目前内置了两个reader

HttpReader - HTTP请求
OssReader - 阿里云oss存储请求

## 中间文件

xxx.db - 分片状态文件，如果一次没有下载成功，会根据db文件中的状态继续下载，忽略已经完成的部分。例如已经下载了4个分片，还有2个就继续下载这两个。
xxx.dp0 - 结尾是一个数字从0-n表示分片，除了最后一个分片，每个分片的大小是partSize

## 修改workers

```go
func main() {
    langgo.Run(&download.Instance{})
    download.Get().Tune(10)
}
```

Tune(size int) - 改变协程池里的协程数量