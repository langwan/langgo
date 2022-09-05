package langgo

import (
	"errors"
	"gopkg.in/yaml.v3"
	"os"
)

var componentConfiguration = make(map[string]interface{})

func LoadConfigurationFile(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(data, &componentConfiguration)
	if err != nil {
		return err
	}
	return nil
}

func GetComponentConfiguration(name string, conf interface{}) error {
	if obj, ok := componentConfiguration[name]; ok {
		marshal, err := yaml.Marshal(obj)
		if err != nil {
			return err
		}
		err = yaml.Unmarshal(marshal, conf)
		if err != nil {
			return err
		}
		return nil
	} else {
		return errors.New("component configuration not find")
	}
}
