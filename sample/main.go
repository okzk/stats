package main

import (
	"github.com/okzk/stats"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	t := stats.SchedulePeriodically(time.Minute, func(s *stats.Stats) { log.Println(s) })
	defer t.Stop()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-sigCh
}
