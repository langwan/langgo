package hello

import "github.com/langwan/langgo/core"

type Instance struct {
	Message string `yaml:"message"`
}

const name = "hello"

var instance *Instance

func (i *Instance) Load() error {
	core.GetComponentConfiguration(name, i)
	return i.Run()
}

func (i *Instance) Run() error {
	instance = i
	return nil
}

func (i *Instance) GetName() string {
	return name
}

func GetInstance() *Instance {
	return instance
}

func Get() *Instance {
	return instance
}
