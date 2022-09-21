package sqlite

import (
	"fmt"
	"gorm.io/gorm"
	"log"
)

const name = "mysql"

var instance *Instance

var db *gorm.DB

type Instance struct {
	Dsn string `yaml:"dsn"`
}

func (i *Instance) GetName() string {
	return name
}

func (i *Instance) Load() error {
	dp := fmt.Sprintf("%s/app.db", dir.GetDefaultDocumentFolderPath())

	db, err = gorm.Open(sqlite.Open(dp), &gorm.Config{})
	if err != nil {
		log.Panicln(err)
		return
	}
	db.AutoMigrate(&model2.Download{}, &model2.FileItem{})
	return nil
}

func (i Instance) Defer() {

}

func (i Instance) Backend() {

}

func Get() *gorm.DB {
	return db
}
