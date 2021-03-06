package os

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// CatchTermination listens for os signals to cancel global context
func CatchTermination(cancel context.CancelFunc) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	log.Print("[WARN] caught termination signal")
	cancel()
}
