package sqlite

import (
	"github.com/langwan/langgo"
	"github.com/langwan/langgo/core"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"os"
	"testing"
)

type account struct {
	gorm.Model
	Name string
}

func TestInstance_Load(t *testing.T) {
	dbPath := "../../testdata/sqlite_test.db"
	os.Remove(dbPath)
	defer os.Remove(dbPath)
	core.EnvName = core.Development
	core.LoadConfigurationFile("../../testdata/configuration_test.app.yml")
	langgo.Run(&Instance{})
	err := Get().AutoMigrate(&account{})
	assert.NoError(t, err)
	acc := account{Name: "langgo"}
	Get().Create(&acc)
	acc2 := account{}
	Get().First(&acc2, "name=?", acc.Name)
	assert.Equal(t, acc.Name, acc2.Name)

}

func TestInstance_Run(t *testing.T) {
	dbPath := "../../testdata/sqlite_test.db"
	os.Remove(dbPath)
	defer os.Remove(dbPath)
	langgo.LoadComponents(&Instance{Path: dbPath})
	err := Get().AutoMigrate(&account{})
	assert.NoError(t, err)
	acc := account{Name: "langgo"}
	Get().Create(&acc)
	acc2 := account{}
	Get().First(&acc2, "name=?", acc.Name)
	assert.Equal(t, acc.Name, acc2.Name)
}
