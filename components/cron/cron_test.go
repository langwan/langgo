package cron

import (
	"fmt"
	"github.com/langwan/langgo"
	"gopkg.in/yaml.v3"
	"os"
	"testing"
	"time"
)

type MyJob struct {
	Name string
}

func (j MyJob) Run() {
	fmt.Println(j.Name, time.Now())
}

func TestBasic(t *testing.T) {
	langgo.Run(&Instance{WithSeconds: true})
	wait := make(chan struct{})
	Get().BindTaskAndSchedule("basic", "* * * * * *", MyJob{Name: "basic"})
	<-wait
}

func TestLoad(t *testing.T) {
	langgo.Run(&Instance{WithSeconds: true})
	wait := make(chan struct{})
	Get().BindTask("my job 1", MyJob{Name: "my job 1"})
	Get().BindTask("my job 2", MyJob{Name: "my job 2"})
	data, err := os.ReadFile("../../testdata/cron.yml")
	if err != nil {
		return
	}
	var schedules []Schedule
	yaml.Unmarshal(data, &schedules)
	Get().Load(schedules...)
	for _, entry := range GetCron().Entries() {
		fmt.Println("entry", entry)
	}
	<-wait
}

func TestWithSecondsFalse(t *testing.T) {
	langgo.Run(&Instance{})
	wait := make(chan struct{})
	Get().BindTaskAndSchedule("basic", "* * * * *", MyJob{Name: "basic"})
	<-wait
}
