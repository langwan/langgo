---
title: "progress"
---

# progress

监听传输进度

## 使用

```go
type ProgressListener interface {
	ProgressChanged(event *ProgressEvent)
}
```
在 download 组件中使用了 ProgressListener 接口，用来观察下载进度。

```go
// ProgressEvent defines progress event
type ProgressEvent struct {
	ConsumedBytes int64
	TotalBytes    int64
	RwBytes       int64
	EventType     ProgressEventType
}
```

ConsumedBytes - 已经传输的字节数

TotalBytes - 需要传输的总字节数

RwBytes - 目前等同于ConsumedBytes

EventType - 事件类型