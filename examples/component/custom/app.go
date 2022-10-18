package main

import (
	"fmt"
	"github.com/langwan/langgo"
	"github.com/langwan/langgo/examples/component/custom/my"
)

func main() {
	langgo.Run(&my.Instance{})
	fmt.Printf("component name is `%s`, message is `%s`\n", my.Get().GetName(), my.Get().Message)
}
