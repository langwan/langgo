package cron

import (
	"errors"
	rcron "github.com/robfig/cron/v3"
)

var c *rcron.Cron

var jobs = make(map[string]*Job)

type Job struct {
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Id          rcron.EntryID `json:"id"`
	Job         rcron.Job     `json:"job"`
}

type Instance struct {
}

const name = "cron"

var instance *Instance

func (i *Instance) Run() error {
	instance = i
	c = rcron.New()
	c.Start()
	return nil
}

func (i *Instance) GetName() string {
	return name
}

func Get() *rcron.Cron {
	return c
}

type Schedule struct {
	Name string `json:"name"`
	Spec string `json:"spec"`
}

func Load(schedules ...Schedule) (errMap map[string]error) {
	errMap = make(map[string]error)
	for _, schedule := range schedules {
		if job, ok := jobs[schedule.Name]; ok {
			var err error
			job.Id, err = c.AddJob(schedule.Spec, job.Job)
			if err != nil {
				errMap[schedule.Name] = err
			}
		}
	}
	return errMap
}

func AddJobAndSchedule(schedule *Schedule, job rcron.Job) error {
	if _, ok := jobs[schedule.Name]; ok {
		return errors.New("job exists")
	}
	jobs[schedule.Name] = &Job{
		Id:  0,
		Job: job,
	}
	return AddSchedule(schedule)
}

func AddJob(job *Job) error {
	if _, ok := jobs[job.Name]; ok {
		return errors.New("job exists")
	}
	jobs[job.Name] = job
	return nil
}

func UpdateSchedule(schedule *Schedule) error {
	if job, ok := jobs[schedule.Name]; ok {
		c.Remove(job.Id)
		var err error
		job.Id, err = c.AddJob(schedule.Spec, job.Job)
		return err
	} else {
		return errors.New("job exists")
	}
	return nil
}

func AddSchedule(schedule *Schedule) error {
	if job, ok := jobs[name]; ok {
		var err error
		job.Id, err = c.AddJob(schedule.Spec, job.Job)
		if err != nil {
			return err
		}
		return nil
	} else {
		return errors.New("job not exists")
	}
}
