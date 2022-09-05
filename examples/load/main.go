package main

import (
	"fmt"
	_ "github.com/langwan/langgo"
	"github.com/langwan/langgo/components/jwt"
	"github.com/langwan/langgo/core"
)

func main() {
	core.AddComponents(&jwt.Instance{})
	core.LoadComponents()
	sign, err := jwt.Sign("langgo")
	if err != nil {
		panic(err)
	}
	err = jwt.Verify(sign)
	if err != nil {
		panic(err)

	}
	fmt.Println("ok")
}
