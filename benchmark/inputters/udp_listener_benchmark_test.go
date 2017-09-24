package benchmark

import (
	"github.com/jadekler/git-go-scalability-talk/benchmark"
	"net"
	"testing"
)

func BenchmarkUdpListener(b *testing.B) {
	for i := 0; i < b.N; i++ {
		u.wg.Add(1)
		sendUdpRequest(u.p, u.laddr, u.raddr)
	}

	u.wg.Wait()
}

func BenchmarkUdpListenerParallel(b *testing.B) {
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            u.wg.Add(1)
            sendUdpRequest(u.p, u.laddr, u.raddr)
        }
    })

    u.wg.Wait()
}

func sendUdpRequest(port int, laddr, raddr *net.UDPAddr) {
	conn, err := net.DialUDP("udp", laddr, raddr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	conn.Write([]byte(benchmark.SMALL_MESSAGE))
}
