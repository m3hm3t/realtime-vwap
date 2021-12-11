package realtime_vwap

import (
	"context"
	"github.com/m3hm3t/realtime-vwap/internal/ws"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// @title Realtime VWAP
// @version 1.0.0
// @description This is a Realtime VWAP CLI App
func main() {

	ctx := context.Background()

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

		s := <-quit

		log.Printf("received signal: %s", s)

		os.Exit(0)
	}()

	err := ws.Run(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
