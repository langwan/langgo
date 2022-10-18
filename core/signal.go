package core

import (
	"fmt"
	"os"
	"os/signal"
)

var handlers = make(map[os.Signal][]func(sig os.Signal))

func SignalHandlers(handler func(sig os.Signal), signals ...os.Signal) {
	for _, s := range signals {
		handlers[s] = append(handlers[s], handler)
	}
}

var SignalNotifyIsRunning = false

func SignalNotify() {
	if SignalNotifyIsRunning {
		if EnvName == Development {
			fmt.Println("SignalNotify can be executed only once")
			return
		}
	}
	SignalNotifyIsRunning = true
	c := make(chan os.Signal)

	var signals []os.Signal

	for sig, _ := range handlers {
		signals = append(signals, sig)
	}

	if len(signals) > 0 {
		signal.Notify(c, signals...)
		go func() {
			for {
				s := <-c
				if hs, ok := handlers[s]; ok {
					for _, handler := range hs {

						handler(s)

					}
				}
			}
		}()
	}

}
