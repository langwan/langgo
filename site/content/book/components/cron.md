---
title: "cron"
---
# cron 计划任务

cron 是 https://github.com/robfig/cron 库的二次封装，便于进行可视化管理。

## 意义

robfig/cron 已经完成了计划任务的所有基础功能，在实际项目中，我们还有两个需要自己实现的需求。

1. 配置化，不实现后台管理功能，大部分项目也会实现通过配置文件、或者在代码当中组装计划任务，来达到管理的目的。

2. 后台可视化管理，大多数项目都有可视化后台，希望在后台可以自由的管理计划任务。

## 定义

每一个cron由 name spec entryId job 四个元素构成

name - 每一个cron的名称，例如后台显示、作为日志当中的tag，这个name是langgo框架定义的，必须是惟一的。

spec - robfig/cron 当中原始的属性，用来表示cron的策略。

entryId - entry 是 robfig/cron 把当前存在的cron称为entry，每一个entry会有一个id，通过id来删除已经存在的计划任务。

job - job 是计划任务实际的执行实例。

## 基本使用方法 在代码当中组装

```go

import (
    rcron "github.com/robfig/cron/v3"
    "fmt"
)

type MyJob struct {
	Name string
}

func (j MyJob) Run() {
	fmt.Println(j.Name, time.Now())
}

func main() {
    langgo.Run(&cron.Instance{WithSeconds: true})
	wait := make(chan struct{})
	cron.Get().BindTaskAndSchedule("basic", "* * * * * *", MyJob{Name: "basic"})
    <-wait
}
```

MyJob 是 rcron.Job 的具体实例，通过 BindTaskAndSchedule() 方法，对具体的实例 MyJob 关联 name 和 spec 属性。其中 spec 属性是计划任务的执行策略。当WithSeconds=true，策略包含秒，WithSeconds=false，策略不包秒只能执行到分，所以当前的例子是每秒钟执行一次 MyJob.Run() 方法

## WithSeconds

查看 robfig/cron 源代码 当 WithSeconds = false

```
var standardParser = NewParser(
	Minute | Hour | Dom | Month | Dow | Descriptor,
)
```

当 WithSeconds = true

```
func WithSeconds() Option {
	return WithParser(NewParser(
		Second | Minute | Hour | Dom | Month | Dow | Descriptor,
	))
}
```


## 从配置中加载

定义配置文件cron.yml

```yaml
- name: "my job 1"
  spec: "* * * * * *"
- name: "my job 2"
  spec: "*/5 * * * * *"
```

代码：

```go
import (
    rcron "github.com/robfig/cron/v3"
    "gopkg.in/yaml.v3"
    "fmt"
    "os"
)

func main() {
    langgo.Run(&cron.Instance{WithSeconds: true})
	wait := make(chan struct{})
	cron.Get().BindTask("my job 1", MyJob{Name: "my job 1"})
	cron.Get().BindTask("my job 2", MyJob{Name: "my job 2"})
	data, err := os.ReadFile("../../testdata/cron.yml")
	if err != nil {
		return
	}
	var schedules []Schedule
	yaml.Unmarshal(data, &schedules)
	cron.Get().Load(schedules...)
	for _, entry := range cron.Get().cron.Entries() {
		fmt.Println(entry)
	}
	<-wait
}
```
调用 BindTask() 方法给每一个 MyJob 实例绑定一个 name 属性。调用 Load() 方法给每一个 task 绑定一个 spec 属性，完成计划任务的初始化。当前这个例子 my job 1 每秒执行一次，my job 2 5秒执行一次

## 后台UI可视化管理

如果您希望从后台UI管理所有的计划任务，可以按照下面的方式实现。

```go

...

cron.Get().BindTask("my job 1", MyJob{Name: "my job 1"})
cron.Get().BindTask("my job 2", MyJob{Name: "my job 2"})

for _, task := range cron.Get().Tasks() {
    fmt.Println(task)
}

...

```

通过 BindTask 方法在代码中绑定出所有的Task。通过 Tasks() 方法拿到所有的name，封装一个系统API提供给后台UI，实现计划任务的选择列表。

```go

...

entryId, _ := cron.Get().UpdateSchedule("my job 1", "* 1 * * * * *")

...

```

实现一个API，通过调用 UpdateSchedule() 方法刷新cron的spec属性。返回值是cron的 entryId，通过这个entryId可以删除已经存在的计划任务。



```go

...

for _, entry := range cron.GetCron().Entries() {
    fmt.Println("entry", entry)
}

...

```

输出:

```shell
entry {1 0x1400017f140 2022-11-17 14:44:13 +0800 CST 0001-01-01 00:00:00 +0000 UTC 0x102af6a10 {my job 1}}
entry {2 0x1400017f180 2022-11-17 14:44:15 +0800 CST 0001-01-01 00:00:00 +0000 UTC 0x102af6a10 {my job 2}}
```

通过 GetCron().Entries() 方法可以打印出当前生效的所有计划任务，可以编写一个API提供给后台UI展示所有当前计划任务。


```go

...

entryId, _ := cron.Get().RemoveTask("my job 1")

...

```

通过 RemoveTask() 方法移除已经存在的cron