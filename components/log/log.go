package log

import (
	"github.com/langwan/langgo/core"
)

type Instance struct {
}

const name = "log"

var instance *Instance

func (i *Instance) Load() {
	instance = i
	core.GetComponentConfiguration(name, i)
}

func (i *Instance) GetName() string {
	return name
}
