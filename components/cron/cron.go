package cron

import (
	"errors"
	"fmt"
	rcron "github.com/robfig/cron/v3"
)

type Task struct {
	Name string        `json:"name"`
	Id   rcron.EntryID `json:"id"`
	Job  rcron.Job     `json:"job"`
}

type Instance struct {
	tasks       map[string]*Task
	cron        *rcron.Cron
	WithSeconds bool `json:"with_seconds"`
}

const moduleName = "cron"

var instance *Instance

func (i *Instance) Run() error {
	instance = i
	l := Logger{}
	var opts []rcron.Option
	opts = append(opts, rcron.WithChain(rcron.Recover(&l)), rcron.WithLogger(&l))
	if i.WithSeconds {
		opts = append(opts, rcron.WithSeconds())
	}
	i.cron = rcron.New(opts...)
	i.tasks = make(map[string]*Task)
	i.cron.Start()
	return nil
}

func (i *Instance) GetName() string {
	return moduleName
}

func Get() *Instance {
	return instance
}

func GetCron() *rcron.Cron {
	return instance.cron
}

type Schedule struct {
	Name string `json:"moduleName"`
	Spec string `json:"spec"`
}

func (i *Instance) Load(schedules ...Schedule) (errMap map[string]error) {
	errMap = make(map[string]error)
	for _, schedule := range schedules {
		if job, ok := instance.tasks[schedule.Name]; ok {
			var err error
			job.Id, err = instance.cron.AddJob(schedule.Spec, job.Job)
			if err != nil {
				errMap[schedule.Name] = err
			}
		}
	}
	return errMap
}

func (i *Instance) BindTask(name string, job rcron.Job) error {
	if _, ok := instance.tasks[name]; ok {
		return errors.New(fmt.Sprintf("job %s exists", name))
	}
	instance.tasks[name] = &Task{
		Name: name,
		Job:  job,
	}
	return nil
}

func (i *Instance) RemoveTask(name string) {
	if job, ok := instance.tasks[name]; ok {
		if job.Id != 0 {
			instance.cron.Remove(job.Id)
		}
		delete(instance.tasks, moduleName)
	}
}

func (i *Instance) BindSchedule(name, spec string) error {
	if job, ok := instance.tasks[name]; ok {
		var err error
		job.Id, err = instance.cron.AddJob(spec, job.Job)
		if err != nil {
			return err
		}
		return nil
	} else {
		return errors.New(fmt.Sprintf("job %s not exists", name))
	}
}

func (i *Instance) UpdateSchedule(name, spec string) (rcron.EntryID, error) {
	if job, ok := instance.tasks[name]; ok {
		if job.Id != 0 {
			instance.cron.Remove(job.Id)
		}
		var err error
		job.Id, err = instance.cron.AddJob(spec, job.Job)
		return job.Id, err
	} else {
		return job.Id, errors.New(fmt.Sprintf("job %s not exists", name))
	}
}

func (i *Instance) BindTaskAndSchedule(name, spec string, job rcron.Job) error {
	if _, ok := instance.tasks[name]; ok {
		return errors.New(fmt.Sprintf("job %s exists", name))
	}
	instance.tasks[name] = &Task{
		Job: job,
	}
	return i.BindSchedule(name, spec)
}

func (i *Instance) Tasks() map[string]*Task {
	return instance.tasks
}
