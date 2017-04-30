package benchmark

import (
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

func sendUdpRequest(port int) {
	raddr := &net.UDPAddr{IP: []byte{byte(127), byte(0), byte(0), byte(1)}, Port: port, Zone: "eth0"}
	conn, err := net.DialUDP("udp", nil, raddr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	conn.Write([]byte(benchmark.SMALL_MESSAGE))
}
