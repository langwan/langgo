package hello

import (
	"github.com/langwan/langgo"
	"github.com/langwan/langgo/core"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInstance_Load(t *testing.T) {
	core.EnvName = core.Development
	core.LoadConfigurationFile("../../testdata/configuration_test.app.yml")
	langgo.Run(&Instance{})
	assert.Equal(t, Get().Message, "hello")
}

func TestInstance_Run(t *testing.T) {
	langgo.Run(&Instance{Message: "hello"})
	assert.Equal(t, Get().Message, "hello")
}
