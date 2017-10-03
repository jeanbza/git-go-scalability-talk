// Used code from: https://varshneyabhi.wordpress.com/2014/12/23/simple-udp-clientserver-in-golang/
package listeners

import (
	"bytes"
	"fmt"
	"github.com/jadekler/git-go-scalability-talk/application/queues"
	"io"
	"net"
)

type TcpListener struct {
	port int
}

func NewTcpListener(port int) *TcpListener {
	return &TcpListener{port: port}
}

func (l *TcpListener) StartAccepting(q queues.Queue) {
	fmt.Printf("Starting TCP listening on port %d\n", l.port)

	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", l.port))
	if err != nil {
		panic(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}

		go readFromConn(conn, q) // incl process buffer manually, retry logic, etc
	}
}

func readFromConn(c net.Conn, q queues.Queue) {
	var add = make(chan ([]byte), 1024)

	go processBuffer(add, q)

	for {
		msg := make([]byte, 1024)
		_, err := c.Read(msg)

		if err != nil {
			if err == io.EOF {
				c.Close()
				return
			}

			panic(err)
		}

		add <- msg
	}
}

func processBuffer(add chan ([]byte), q queues.Queue) {
	b := bytes.NewBuffer([]byte{})

	for {
		select {
		case msg := <-add:
			b.Write(msg)
			break
		default:
			l, err := b.ReadBytes('\n')
			if err != nil {
				if err == io.EOF {
					break
				}

				panic(err)
			}

			if len(l) <= 1 || l[0] == 0 { // kinda hacky way to check if it's a fully-formed message
				break
			}

			q.Enqueue(l)
		}
	}
}
