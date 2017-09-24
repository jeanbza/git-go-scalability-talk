package benchmark

import (
	"github.com/jadekler/git-go-scalability-talk/benchmark"
	"net"
	"testing"
)

func BenchmarkUdpListener(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sendUdpRequest(u.conn)
	}

	<- u.t.C
}

func BenchmarkUdpListenerParallel(b *testing.B) {
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            sendUdpRequest(u.conn)
        }
    })

	<- u.t.C
}

func sendUdpRequest(conn net.Conn) {
	conn.Write([]byte(benchmark.SMALL_MESSAGE))
}
