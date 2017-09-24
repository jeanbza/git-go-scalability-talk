package benchmark

import (
	"net"
	"testing"
)

func BenchmarkStreamingUdpListener(b *testing.B) {
	for i := 0; i < b.N; i++ {
		streamUdpRequest(su.conn)
	}

	<-su.t.C
}

func BenchmarkStreamingUdpListenerParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			streamUdpRequest(su.conn)
		}
	})

	<-su.t.C
}

func streamUdpRequest(conn net.Conn) {
	conn.Write([]byte(msg))
}
