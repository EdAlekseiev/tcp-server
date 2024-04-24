package main

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/EdAlekseiev/tcp-server/internal/transport/tcp"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	server := tcp.NewTcpTransport(ctx, "localhost:8085")

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := server.ListenAndAccept(); err != nil {
			panic(err)
		}
	}()
	<-signalChan
	cancel()
	wg.Wait()
}
