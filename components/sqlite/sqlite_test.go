package sqlite

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"os"
	"testing"
)

func TestSqlite(t *testing.T) {
	type Account struct {
		gorm.Model
		Name string
	}
	ins := Instance{}
	ins.Path = "../../testdata/sqlite_test.db"
	ins.load()
	os.Remove(instance.Path)
	instance.Load()
	Get().AutoMigrate(&Account{})
	acc := Account{Name: "langgo"}
	Get().Create(&acc)
	acc2 := Account{}
	Get().First(&acc2, "name=?", acc.Name)
	assert.Equal(t, acc.Name, acc2.Name)
	os.Remove(instance.Path)
}
