package main

import (
	"github.com/langwan/langgo"
	"github.com/langwan/langgo/components/mysql"
	"github.com/langwan/langgo/core/log"
)

func main() {
	langgo.Run(&mysql.Instance{})
	mysql.Main().AutoMigrate(&Account{})
	acc := Account{
		Name: "chihuo",
	}
	mysql.Main().Create(&acc)
	acc.Name = "famingjia"
	mysql.Main().Save(&acc)
	newAcc := Account{}
	mysql.Main().First(&newAcc, "id=?", acc.ID)
	log.Logger("app", "main").Info().Interface("newAcc", newAcc).Send()
	mysql.Main().Unscoped().Delete(&Account{}, newAcc.ID)
}
