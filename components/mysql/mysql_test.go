package mysql

import (
	"github.com/langwan/langgo/core"
	"github.com/langwan/langgo/core/log"
	"testing"
)

func Test_Mysql(t *testing.T) {
	core.EnvName = core.Development
	core.LoadConfigurationFile("../../testdata/configuration_test.app.yml")
	l := log.Instance{}
	l.Load()
	i := Instance{}
	i.Load()
	var one int
	res := Get().Debug().Select(" 1").Scan(&one)
	if res.RowsAffected > 0 {
		t.Log(one)
	}
}
