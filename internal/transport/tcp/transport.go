package tcp

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/EdAlekseiev/tcp-server/internal/transport"
)

type TcpTransport struct {
	listenAddress string
	listener      net.Listener

	wg  sync.WaitGroup
	ctx context.Context
}

func NewTcpTransport(ctx context.Context, listenAddress string) transport.Transport {
	return &TcpTransport{
		listenAddress: listenAddress,
		ctx:           ctx,
	}
}

func (tcp *TcpTransport) ListenAndAccept() error {
	var err error

	tcp.listener, err = net.Listen("tcp", tcp.listenAddress)
	if err != nil {
		panic(err)
	}

	tcp.wg.Add(1)
	go func() {
		defer tcp.wg.Done()

		tcp.startAcceptLoop()
	}()

	fmt.Printf("TCP: listening address %s\n", tcp.listenAddress)
	<-tcp.ctx.Done()
	tcp.listener.Close()
	tcp.wg.Wait()
	return nil
}

func (tcp *TcpTransport) startAcceptLoop() {
	for {
		conn, err := tcp.listener.Accept()
		if err != nil {
			if errors.Is(err, net.ErrClosed) {
				return
			}

			fmt.Printf("TCP: error during accepr connection %v\n", err)
		}

		tcp.wg.Add(1)
		go func() {
			defer tcp.wg.Done()

			tcp.handleConnection(conn)
		}()
	}
}

func (tcp *TcpTransport) handleConnection(conn net.Conn) {
	defer func() {
		defer conn.Close()
		fmt.Printf("TCP: Connection was closed\n")
	}()

	for {
		select {
		case <-tcp.ctx.Done():
			return
		default:
			conn.SetDeadline(time.Now().Add(1 * time.Second))
			msg, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil {
				if opErr, ok := err.(*net.OpError); ok && opErr.Timeout() {
					continue
				}
				fmt.Printf("TCP: Can't read message. Err: %v\n", err)
				return
			}

			fmt.Println("-> ", msg)
		}
	}
}
