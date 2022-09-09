package main

import (
	"github.com/langwan/langgo"
	_ "github.com/langwan/langgo"
	"github.com/langwan/langgo/components/jwt"
	"github.com/langwan/langgo/core/log"
)

func main() {
	langgo.Run(&jwt.Instance{})
	sign, err := jwt.Sign("langgo")
	if err != nil {
		panic(err)
	}
	err = jwt.Verify(sign)
	if err != nil {
		panic(err)

	}
	log.Logger("app", "jwt").Info().Msg("ok")
}
