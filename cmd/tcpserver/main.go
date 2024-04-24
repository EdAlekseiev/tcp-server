package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/EdAlekseiev/tcp-server/internal/transport/tcp"
)

var port *int

func init() {
	port = flag.Int("port", 8085, "port number")
}

func main() {
	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())
	server := tcp.NewTcpTransport(ctx, fmt.Sprintf(":%d", *port))

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
