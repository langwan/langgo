package ffmpeg

import "time"

type Instance struct {
	FFMpeg         string        `json:"ffmpeg"`
	FFProbe        string        `json:"ffprobe"`
	CommandTimeout time.Duration `json:"command_timeout"`
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
