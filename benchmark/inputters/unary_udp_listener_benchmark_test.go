package benchmark

import (
	"net"
	"testing"
)

func BenchmarkUnaryUdpListener(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sendUdpRequest(uu.laddr, uu.raddr)
	}

	<-uu.t.C
}

func BenchmarkUnaryUdpListenerParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			sendUdpRequest(uu.laddr, uu.raddr)
		}
	})

	<-uu.t.C
}

func sendUdpRequest(laddr, raddr *net.UDPAddr) {
	conn, err := net.DialUDP("udp", laddr, raddr)
	if err != nil {
		panic(err)
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}()

	conn.Write([]byte(msg))
}
