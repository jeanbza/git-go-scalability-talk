package benchmark

import (
	"fmt"
	"github.com/jadekler/git-go-scalability-talk/application/inputters"
	"github.com/jadekler/git-go-scalability-talk/application/queues"
	"github.com/jadekler/git-go-scalability-talk/benchmark"
	"net"
	"sync"
	"testing"
)

var u udpListenerBenchmark = udpListenerBenchmark{}

type udpListenerBenchmark struct {
	l  *listeners.UdpListener
	wg *sync.WaitGroup
	q  queues.Queue
	p  int
}

func BenchmarkUdpListener(b *testing.B) {
	if u.l == nil {
		u.p = benchmark.GetOpenTcpPort()

		u.wg = &sync.WaitGroup{}
		u.q = benchmark.NewWaitingQueue(u.wg)

		u.l = listeners.NewUdpListener(u.p)
		go u.l.StartAccepting(u.q)
	}

	for i := 0; i < b.N; i++ {
		u.wg.Add(1)
		sendUdpRequest(u.p)
	}

	u.wg.Wait()
}

func BenchmarkUdpListenerParallel(b *testing.B) {
    if u.l == nil {
        u.p = benchmark.GetOpenTcpPort()

        u.wg = &sync.WaitGroup{}
        u.q = benchmark.NewWaitingQueue(u.wg)

        u.l = listeners.NewUdpListener(u.p)
        go u.l.StartAccepting(u.q)
    }

    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            u.wg.Add(1)
            sendUdpRequest(u.p)
        }
    })

    u.wg.Wait()
}

func sendUdpRequest(port int) {
	raddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("localhost:%d", port))
	laddr, err := net.ResolveUDPAddr("udp", "localhost:0")
	if err != nil {
		panic(err)
	}
	conn, err := net.DialUDP("udp", laddr, raddr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	conn.Write([]byte(benchmark.MEDIUM_MESSAGE))
}
