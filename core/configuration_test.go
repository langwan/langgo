package core

import "testing"

func TestLoadConfigurationFile(t *testing.T) {
	err := LoadConfigurationFile("../testdata/configuration_test.app.yml")
	if err != nil {
		t.Error(err)
		return
	}
	if conf, ok := componentConfiguration["jwt"]; ok {
		if name, ok := conf.(map[string]interface{})["Secret"]; ok {
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
	err := LoadConfigurationFile("../testdata/configuration_test.app.yml")
	if err != nil {
		t.Error(err)
		return
	}
	conf := struct {
		Secret string `yaml:"secret"`
		Level  int    `yaml:"level"`
	}{Level: 1}
	err = GetComponentConfiguration("jwt", &conf)
	if err != nil {
		t.Error(err)
	} else {
		if conf.Secret == "123456" && conf.Level == 1 {

		} else {
			t.Fatal("name error")
		}
	}
}
