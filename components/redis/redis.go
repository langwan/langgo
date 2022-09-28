package redis

import (
	"github.com/go-redis/redis"
	"github.com/langwan/langgo/core"
)

const name = "redis"

var instance *Instance

type item struct {
	Dsn string `yaml:"dsn"`
}

var connections = make(map[string]*redis.Client)

type Instance struct {
}

func (i *Instance) GetName() string {
	return name
}

func (i *Instance) Load() error {
	instance = i

	items := make(map[string]item)

	core.GetComponentConfiguration(name, &items)

	for k, c := range items {
		opt, err := redis.ParseURL(c.Dsn)
		if err != nil {
			return err
		}
		connections[k] = redis.NewClient(opt)
	}
	return nil
}

func Main() *redis.Client {
	return Get("main")
}
func Get(name string) *redis.Client {
	conn, ok := connections[name]
	if ok {
		return conn
	} else {
		return nil
	}
}
