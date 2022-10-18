package sqlite

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const name = "sqlite"

var db *gorm.DB

type Instance struct {
	Path string `yaml:"path"`
}

var instance *Instance

func (i *Instance) GetName() string {
	return name
}

func (i *Instance) Run() error {
	instance = i
	var err error
	db, err = gorm.Open(sqlite.Open(i.Path), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}

func Get() *gorm.DB {
	return db
}
