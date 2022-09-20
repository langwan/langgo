package mysql

import (
	"github.com/langwan/langgo/core"
	"github.com/langwan/langgo/core/log"
	gormMysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

const name = "mysql"

var instance *Instance

type item struct {
	Dsn             string        `yaml:"dsn"`
	MaxIdleConns    int           `yaml:"max_idle_conns"`
	MaxOpenConns    int           `yaml:"max_open_conns"`
	ConnMaxLifetime time.Duration `yaml:"conn_max_lifetime"`
}

var connections = make(map[string]*gorm.DB)

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
		conn, err := gorm.Open(gormMysql.Open(c.Dsn), &gorm.Config{Logger: NewWithLogger(*log.Logger("mysql", "info")), SkipDefaultTransaction: true})
		if err != nil {
			log.Logger("component", "mysql").Warn().Err(err).Send()
			continue
		}

		sqlDB, err := conn.DB()
		if err != nil {
			log.Logger("component", "mysql").Warn().Err(err).Send()
			continue
		}

		sqlDB.SetMaxIdleConns(c.MaxOpenConns)
		sqlDB.SetMaxOpenConns(c.MaxOpenConns)
		sqlDB.SetConnMaxLifetime(c.ConnMaxLifetime)

		connections[k] = conn
	}
	return nil
}

func Main() *gorm.DB {
	return Get("main")
}
func Get(name string) *gorm.DB {
	conn, ok := connections[name]
	if ok {
		return conn
	} else {
		return nil
	}
}
