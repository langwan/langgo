package core

type Component interface {
	Run() error
	GetName() string
}

var components = make(map[string]Component)

func LoadComponents(instances ...Component) {
	for _, c := range instances {
		components[c.GetName()] = c
		GetComponentConfiguration(c.GetName(), c)
		c.Run()
	}
}

// LoadComponent load component
func LoadComponent(c Component) {
	GetComponentConfiguration(c.GetName(), c)
	c.Run()
}

func RunComponents(instances ...Component) {
	for _, c := range instances {
		components[c.GetName()] = c
		c.Run()
	}
}
