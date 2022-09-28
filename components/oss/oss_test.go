package oss

import (
	"github.com/langwan/langgo/core"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOss(t *testing.T) {
	core.EnvName = core.Development
	core.LoadConfigurationFile("../../testdata/configuration_test.app.yml")
	i := Instance{}
	i.Load()
	object, err := GetObject("232a78e09f2c441899c5b90d4333bf2e/b414e97a355b4e60b6e603d8cee21034/未命名.png")
	assert.NoError(t, err)
	assert.NotNil(t, object)
}
