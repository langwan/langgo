package main

import (
	"fmt"
	"github.com/langwan/langgo"
	_ "github.com/langwan/langgo"
	"github.com/langwan/langgo/components/jwt"
	"github.com/langwan/langgo/components/log"
	"github.com/langwan/langgo/core"
	"syscall"
)

func main() {
	die := make(chan bool)
	core.AddComponents(&jwt.Instance{})
	core.LoadComponents()
	core.SignalHandle(&core.SignalHandler{
		Sig: syscall.SIGALRM,
		F: func() {
			die <- true
		},
	})
	langgo.Run()
	sign, err := jwt.Sign("langgo")
	if err != nil {
		panic(err)
	}
	err = jwt.Verify(sign)
	if err != nil {
		panic(err)

	}
	log.Logger("app", "jwt").Info().Msg("ok")
	<-die
	log.Logger("sys", "app").Info().Msg("exit")
	fmt.Println("exit")
}
