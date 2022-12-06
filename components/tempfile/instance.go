package tempfile

type Instance struct {
	Base string `yaml:"base"`
	inst *TempFile
}

const name = "tempfile"

var instance *Instance

func (i *Instance) Run() error {
	instance = i
	instance.inst = &TempFile{Base: instance.Base}
	return nil
}

func (i *Instance) GetName() string {
	return name
}

func Get() *TempFile {
	return instance.inst
}
