package password

import (
	"crypto/sha1"
	"encoding/hex"
	"github.com/langwan/langgo/core"
)

type Instance struct {
	Salt string `yaml:"salt"`
}

var instance *Instance

const name = "password"

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

func Hash(orig string) string {
	hn := sha1.New()
	hn.Write([]byte(orig))
	hn.Write([]byte(instance.Salt))
	data := hn.Sum([]byte(""))
	return hex.EncodeToString(data)
}
