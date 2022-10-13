package mysql

import (
	"github.com/langwan/langgo/core"
	"github.com/langwan/langgo/core/log"
	gormMysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"time"
)

const name = "mysql"

var instance *Instance

type item struct {
	Dsn                       string              `yaml:"dsn"`
	MaxIdleConns              int                 `yaml:"max_idle_conns"`
	MaxOpenConns              int                 `yaml:"max_open_conns"`
	ConnMaxLifetime           time.Duration       `yaml:"conn_max_lifetime"`
	SlowThreshold             time.Duration       `yaml:"slow_threshold"`
	Colorful                  bool                `yaml:"colorful"`
	IgnoreRecordNotFoundError bool                `yaml:"ignore_record_not_found_error"`
	LogLevel                  gormlogger.LogLevel `yaml:"log_level"`
}

var connections = make(map[string]*gorm.DB)

type Instance struct {
	items map[string]item
}

func (i *Instance) GetName() string {
	return name
}

func (i *Instance) Load() error {
	i.items = make(map[string]item)
	core.GetComponentConfiguration(name, &i.items)
	return i.Run()
}

func (i *Instance) Run() error {
	instance = i
	for k, c := range i.items {
		zl := log.Logger("mysql", k)

		l := New(*zl, gormlogger.Config{
			SlowThreshold:             c.SlowThreshold,
			Colorful:                  c.Colorful,
			IgnoreRecordNotFoundError: c.IgnoreRecordNotFoundError,
			LogLevel:                  c.LogLevel,
		})

		conn, err := gorm.Open(gormMysql.Open(c.Dsn), &gorm.Config{Logger: l, SkipDefaultTransaction: true})
		if err != nil {
			log.Logger("component", "mysql").Warn().Err(err).Send()
			continue
		}

		sqlDB, err := conn.DB()
		if err != nil {
			log.Logger("component", "mysql").Warn().Err(err).Send()
			continue
		}

		sqlDB.SetMaxIdleConns(c.MaxIdleConns)
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
