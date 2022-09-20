package main

import (
	"github.com/langwan/langgo"
	"github.com/langwan/langgo/components/mysql"
	"github.com/langwan/langgo/core/log"
)

func main() {
	langgo.Run(&mysql.Instance{})
	mysql.Get().AutoMigrate(&Account{})
	acc := Account{
		Name: "chihuo",
	}
	mysql.Get().Create(&acc)
	acc.Name = "famingjia"
	mysql.Get().Save(&acc)
	newAcc := Account{}
	mysql.Get().First(&newAcc, "id=?", acc.ID)
	log.Logger("app", "main").Info().Interface("newAcc", newAcc).Send()
	mysql.Get().Unscoped().Delete(&Account{}, newAcc.ID)
}
