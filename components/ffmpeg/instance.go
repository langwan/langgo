package ffmpeg

import "time"

type Instance struct {
	FFMpeg         string        `yaml:"ffmpeg"`
	FFProbe        string        `yaml:"ffprobe"`
	CommandTimeout time.Duration `yaml:"command_timeout"`
	bind           *Bind
}

const name = "ffmpeg"

var instance *Instance

func (i *Instance) Run() error {
	instance = i
	instance.bind = &Bind{
		FFMpeg:         instance.FFMpeg,
		FFProbe:        instance.FFProbe,
		CommandTimeout: instance.CommandTimeout,
	}
	return nil
}

func (i *Instance) GetName() string {
	return name
}

func Get() *Bind {
	return instance.bind
}
