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
	j := MyJob{"one"}
	AddJob("one", j)
	Load(Schedule{
		Name: "one",
		Spec: "* * * * *",
	})

	j2 := MyJob{"two"}
	AddJob("two", j2)
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
