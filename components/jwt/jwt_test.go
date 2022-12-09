package jwt

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
	payload := struct {
		Name string
	}{Name: "langwan"}
	sign, err := Sign(payload)
	assert.NoError(t, err)
	err = Verify(sign)
	assert.NoError(t, err)
}

func TestInstance_Run(t *testing.T) {
	langgo.LoadComponents(&Instance{Secret: "123456"})
	payload := struct {
		Name string
	}{Name: "langwan"}
	sign, err := Sign(payload)
	assert.NoError(t, err)
	err = Verify(sign)
	assert.NoError(t, err)
}
