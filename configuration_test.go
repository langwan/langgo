package langgo

import "testing"

func TestLoadConfigurationFile(t *testing.T) {
	err := LoadConfigurationFile("./testdata/configuration_test.app.yml")
	if err != nil {
		t.Error(err)
		return
	}
	if conf, ok := componentConfiguration["test"]; ok {
		if name, ok := conf.(map[string]interface{})["name"]; ok {
			if name.(string) != "langgo" {
				t.Fatal("conf name error")
			}
		} else {
			t.Fatal("conf type error")
		}
	} else {
		t.Fatal("test config not find")
	}
}

func TestGetComponentConfiguration(t *testing.T) {
	err := LoadConfigurationFile("./testdata/configuration_test.app.yml")
	if err != nil {
		t.Error(err)
		return
	}
	conf := struct {
		Name  string `yaml:"name"`
		Level int    `yaml:"level"`
	}{Level: 1}
	err = GetComponentConfiguration("test", &conf)
	if err != nil {
		t.Error(err)
	} else {
		if conf.Name == "langgo" && conf.Level == 1 {

		} else {
			t.Fatal("name error")
		}
	}
}
