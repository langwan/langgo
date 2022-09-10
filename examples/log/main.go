package main

import (
	"github.com/langwan/langgo"
	"github.com/langwan/langgo/components/hello"
	"github.com/langwan/langgo/core/log"
	"time"
)

func main() {
	langgo.Run(&hello.Instance{Message: "hello component"})
	log.Logger("component", "hello").Info().Msg(hello.GetInstance().Message)
	//log.Logger("chihuo", "hello").Info().Msg("message")
	loop()
}

func loop() {
	i := 0
	for {
		log.Logger("app", "sleep").Info().Int("index", i).Send()
		i++
		time.Sleep(500 * time.Millisecond)
	}
}
