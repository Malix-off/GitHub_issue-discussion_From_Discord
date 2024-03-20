package keepAlive

import (
	"os"
	"os/signal"
	"syscall"
)

func Main() {
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-s
}
