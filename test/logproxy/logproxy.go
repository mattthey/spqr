package logproxy

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/jackc/pgx/v5/pgproto3"
)

const failedToReceiveMessage = "failed to received msg %w"

func getC() (net.Conn, error) {
	const proto = "tcp"
	const addr = "[::1]:6432"
	return net.Dial(proto, addr)
}

type Proxy struct {
}

func (p *Proxy) Run() error {
	ctx := context.Background()

	listener, err := net.Listen("tcp6", "[::1]:5433")
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer func() {
		if err := listener.Close(); err != nil {
			log.Printf("error closing listener: %v", err)
		}
	}()

	cChan := make(chan net.Conn)

	accept := func(l net.Listener, cChan chan net.Conn) {
		for {
			c, err := l.Accept()
			if err != nil {
				// handle error (and then for example indicate acceptor is down)
				cChan <- nil
				return
			}
			cChan <- c
		}
	}

	go accept(listener, cChan)

	for {
		select {
		case <-ctx.Done():
			os.Exit(1)
		case c := <-cChan:

			go func() {
				log.Fatal(p.serv(c))
			}()
		}
	}
}

func (p *Proxy) serv(netconn net.Conn) error {

	conn, err := getC()
	if err != nil {
		return err
	}

	frontend := pgproto3.NewFrontend(conn, conn)
	cl := pgproto3.NewBackend(netconn, netconn)

	cb := func(msg pgproto3.FrontendMessage) {
		log.Printf("received msg %v", msg)

		switch v := msg.(type) {
		case *pgproto3.Parse:
			log.Printf("received prep stmt %v %v", v.Name, v.Query)
		case *pgproto3.Query:
			log.Printf("received message %v", v.String)
		default:
		}
	}
	shouldStop := func(msg pgproto3.BackendMessage) bool {
		log.Printf("received msg %v", msg)

		switch msg.(type) {
		case *pgproto3.ReadyForQuery:
			return false
		default:
			return false
		}
	}

	for {
		msg, err := cl.Receive()

		cb(msg)

		if err != nil {
			return fmt.Errorf(failedToReceiveMessage, err)
		}
		frontend.Send(msg)
		if err := frontend.Flush(); err != nil {
			return fmt.Errorf(failedToReceiveMessage, err)
		}
		for {
			retmsg, err := frontend.Receive()
			if err != nil {
				return fmt.Errorf(failedToReceiveMessage, err)
			}

			cl.Send(retmsg)
			if err := cl.Flush(); err != nil {
				return fmt.Errorf(failedToReceiveMessage, err)
			}

			if shouldStop(retmsg) {
				break
			}
		}
	}
}
