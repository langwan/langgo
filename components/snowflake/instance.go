package snowflake

type Instance struct {
	MachineID int64 `yaml:"machine_id"`
}

const name = "snowflake"

var instance *Snowflake

func (i *Instance) Run() (err error) {
	instance, err = New(i.MachineID)
	if err != nil {
		return err
	}
	return nil
}

func (i *Instance) GetName() string {
	return name
}

func Get() *Snowflake {
	return instance
}

func Gen() int64 {
	if instance != nil {
		return instance.Gen()
	} else {
		return -1
	}
}
