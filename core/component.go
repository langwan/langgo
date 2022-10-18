package core

type Component interface {
	Run() error
	GetName() string
}

var components = make(map[string]Component)

func LoadComponents() {
	for _, c := range components {
		GetComponentConfiguration(c.GetName(), c)
		c.Run()
	}
}

func AddComponents(instances ...Component) {
	for _, c := range instances {
		components[c.GetName()] = c
	}
}
func RunComponents(instances ...Component) {
	for _, c := range instances {
		components[c.GetName()] = c
		c.Run()
	}
}
