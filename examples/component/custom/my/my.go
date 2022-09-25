package my

import "github.com/langwan/langgo/core"

type Instance struct {
	Message string `yaml:"message"`
}

const name = "my"

var instance *Instance

func (i *Instance) Load() error {
	instance = i
	core.GetComponentConfiguration(name, i)
	return nil
}

func (i *Instance) GetName() string {
	return name
}

func GetInstance() *Instance {
	return instance
}
