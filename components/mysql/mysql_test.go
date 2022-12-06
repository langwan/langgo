package mysql

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
	var one int64
	res := Main().Debug().Raw("SELECT 1").Scan(&one)
	assert.Equal(t, res.RowsAffected, int64(1))
	assert.Equal(t, one, int64(1))
}

func TestInstance_Run(t *testing.T) {
	langgo.LoadComponents(&Instance{Items: map[string]item{"main": {Dsn: "root:123456@tcp(localhost:3306)/simple?charset=utf8mb4&parseTime=True&loc=Local"}}})
	var one int64
	res := Main().Debug().Raw("SELECT 1").Scan(&one)
	assert.Equal(t, res.RowsAffected, int64(1))
	assert.Equal(t, one, int64(1))
}
