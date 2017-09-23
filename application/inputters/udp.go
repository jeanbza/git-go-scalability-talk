// Used code from: https://varshneyabhi.wordpress.com/2014/12/23/simple-udp-clientserver-in-golang/
package listeners

import (
	"fmt"
	"github.com/jadekler/git-go-scalability-talk/application/queues"
	"net"
)

type UdpListener struct {
	port int
}

func NewUdpListener(port int) *UdpListener {
	return &UdpListener{port: port}
}

func (l *UdpListener) StartAccepting(q queues.Queue) {
	fmt.Printf("Starting UDP listening on port %d\n", l.port)

	ServerAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf(":%d", l.port))
	if err != nil {
		fmt.Println(err)
	}

	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	if err != nil {
		fmt.Println(err)
	}
	defer ServerConn.Close()

	buf := make([]byte, 10000)

	for {
		n, _, err := ServerConn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("Error: ", err)
		}

		message := buf[0:n]
		q.Enqueue(message)
	}
}
