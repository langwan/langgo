package cron

import (
	"fmt"
	"github.com/langwan/langgo"
	"testing"
	"time"
)

type MyJob struct {
	Name string
}

func (j MyJob) Run() {
	fmt.Println(j.Name, time.Now())
}

func TestCron(t *testing.T) {
	langgo.Run(&Instance{})
	wait := make(chan struct{})

	AddJob(&Job{
		Name:        "one",
		Description: "one description",
		Id:          0,
		Job:         MyJob{Name: "one"},
	})
	Load(Schedule{
		Name: "one",
		Spec: "* * * * *",
	})

	AddJob(&Job{
		Name:        "two",
		Description: "two description",
		Id:          0,
		Job:         MyJob{Name: "two"},
	})
	Load(Schedule{
		Name: "two",
		Spec: "*/2 * * * *",
	})

	for _, entry := range Get().Entries() {
		t.Log(entry)
	}

	UpdateSchedule(&Schedule{
		Name: "two",
		Spec: "0 17 * * 1-5",
	})
	t.Log("\n")
	for _, entry := range Get().Entries() {
		t.Log(entry)
	}

	<-wait
}

type MyRecoverJob struct {
	Name string
}

func (j MyRecoverJob) Run() {
	panic("MyRecoverJob")
}

func TestRecover(t *testing.T) {
	langgo.Run(&Instance{})
	wait := make(chan struct{})
	AddJobAndSchedule(&Schedule{
		Name: "MyRecoverJob",
		Spec: "* * * * *",
	}, MyRecoverJob{
		Name: "MyRecoverJob",
	})
	<-wait
}
