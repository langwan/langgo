package main

import (
	"fmt"
	"github.com/langwan/langgo"
	"github.com/langwan/langgo/components/mysql"
	"github.com/langwan/langgo/examples/component/custom/my"
)

func main() {
	langgo.Run(&my.Instance{}, &mysql.Instance{})
	fmt.Printf("component name is `%s`, message is `%s`\n", my.GetInstance().GetName(), my.GetInstance().Message)
}
