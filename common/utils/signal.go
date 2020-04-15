package utils

import (
	"os"
	"os/signal"
)

func KillSignal(fn func()) {
	c := make(chan os.Signal, 0)
	signal.Notify(c, os.Interrupt, os.Kill)

	go func() {
		<-c
		fn()
	}()
}
