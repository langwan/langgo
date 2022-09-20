package redis

import (
	"github.com/langwan/langgo/core"
	"github.com/langwan/langgo/core/log"
	"testing"
	"time"
)

func TestRedis(t *testing.T) {
	core.EnvName = core.Development
	core.LoadConfigurationFile("../../testdata/configuration_test.app.yml")
	l := log.Instance{}
	l.Load()
	i := Instance{}
	i.Load()

	cmd := Get().Set("name", "chihuo", 10*time.Second)
	log.Logger("test", "redis").Debug().Interface("cmd", cmd).Send()
	str, err := Get().Get("name").Result()
	log.Logger("test", "redis").Debug().Str("str", str).Err(err).Send()
}
