package core

import (
	"os"
	"os/signal"
	"syscall"
)

type SignalHandler struct {
	Sig syscall.Signal
	F   func()
}

var handlers []*SignalHandler

func SignalHandle(handler *SignalHandler) {
	handlers = append(handlers, handler)
}

func SignalNotify() {
	c := make(chan os.Signal)

	var signals []os.Signal
	for _, handler := range handlers {
		find := false
		for _, sig := range signals {
			if sig == handler.Sig {
				find = true
				break
			}
		}
		if !find {
			signals = append(signals, handler.Sig)
		}
	}

	signal.Notify(c, signals...)
	go func() {
		for {
			s := <-c
			for _, handler := range handlers {
				if handler.Sig == s {
					handler.F()
				}
			}
		}
	}()
}
