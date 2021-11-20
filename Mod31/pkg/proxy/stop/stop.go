package stop

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

func Proxy(cancel context.CancelFunc) {
	go func() {
		exitCh := make(chan os.Signal, 1)
		signal.Notify(exitCh, os.Interrupt, syscall.SIGINT)
		<- exitCh
		cancel()
	}()
}
