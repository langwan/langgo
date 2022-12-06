package redis

import (
	"github.com/langwan/langgo"
	"github.com/langwan/langgo/core"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestInstance_Load(t *testing.T) {
	core.EnvName = core.Development
	core.LoadConfigurationFile("../../testdata/configuration_test.app.yml")
	langgo.Run(&Instance{})
	value := "langgo"
	Main().Set("name", value, 10*time.Second)
	getValue, err := Main().Get("name").Result()
	assert.NoError(t, err)
	assert.Equal(t, value, getValue)
	assert.NotEmpty(t, getValue)
}

func TestInstance_Run(t *testing.T) {
	langgo.LoadComponents(&Instance{Items: map[string]item{"main": {Dsn: "redis://default:redispw@localhost:55000/0"}}})
	value := "langgo"
	Main().Set("name", value, 10*time.Second)
	getValue, err := Main().Get("name").Result()
	assert.NoError(t, err)
	assert.Equal(t, value, getValue)
	assert.NotEmpty(t, getValue)
}
