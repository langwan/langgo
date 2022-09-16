package main

import (
	"github.com/langwan/langgo"
	"github.com/langwan/langgo/components/hello"
	"github.com/langwan/langgo/core/log"
)

func main() {
	langgo.Run(&hello.Instance{Message: "hello component"})
	log.Logger("component", "hello").Info().Msg(hello.GetInstance().Message)
}
