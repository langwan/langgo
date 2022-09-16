package server

import "github.com/langwan/langgo/core"

type Instance struct {
	EtcdHost    string `yaml:"etcd_host"`
	ServiceName string `yaml:"service_name"`
}

const name = "server"

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
