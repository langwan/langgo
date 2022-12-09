package upload

import "github.com/langwan/langgo/helpers/size"

type Instance struct {
	Workers  int    `yaml:"workers"`
	PartSize string `yaml:"part_size"`
	upload   *Upload
}

const name = "upload"

var instance *Instance

func (i *Instance) Run() error {
	instance = i
	partSize, _ := helper_size.RAMInBytes(i.PartSize)
	instance.upload = &Upload{
		Workers:  i.Workers,
		PartSize: partSize,
	}
	instance.upload.Init()
	return nil
}

func (i *Instance) GetName() string {
	return name
}

func Get() *Upload {
	return instance.upload
}
