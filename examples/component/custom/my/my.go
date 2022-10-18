package my

type Instance struct {
	Message string `yaml:"message"`
}

const name = "my"

var instance *Instance

func (i *Instance) Run() error {
	instance = i
	return nil
}

func (i *Instance) GetName() string {
	return name
}

func Get() *Instance {
	return instance
}
