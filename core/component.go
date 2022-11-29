package core

import (
	"gopkg.in/yaml.v3"
)

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

func LoadComponentFromYaml(c Component, content []byte) error {
	err := yaml.Unmarshal(content, c)
	if err != nil {
		return err
	}
	c.Run()
	return nil
}
