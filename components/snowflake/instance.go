package snowflake

type Instance struct {
	nodeId int64 `yaml:"node_id"`
}

const name = "snowflake"

var instance *Node

func (i *Instance) Run() error {
	newNode, err := NewNode(i.nodeId)
	instance = newNode
	if err != nil {
		return err
	}
	return nil
}

func (i *Instance) GetName() string {
	return name
}

func Get() *Node {
	return instance
}
