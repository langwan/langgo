package core

type Component interface {
	Load() error
	GetName() string
}

var components = make(map[string]Component)

func LoadComponents() {

	for _, c := range components {
		c.Load()
	}
}

func AddComponents(instances ...Component) {
	for _, c := range instances {
		components[c.GetName()] = c
	}
}
