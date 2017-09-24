package benchmark

import (
	"github.com/jadekler/git-go-scalability-talk/benchmark"
	"net"
	"testing"
)

func BenchmarkStreamingUdpListener(b *testing.B) {
	for i := 0; i < b.N; i++ {
		streamUdpRequest(u.conn)
	}

	<- u.t.C
}

func BenchmarkStreamingUdpListenerParallel(b *testing.B) {
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            streamUdpRequest(u.conn)
        }
    })

	<- u.t.C
}

func streamUdpRequest(conn net.Conn) {
	conn.Write([]byte(benchmark.SMALL_MESSAGE))
}
