package benchmark

import (
	"fmt"
	"github.com/jadekler/git-go-scalability-talk/application/inputters"
	"github.com/jadekler/git-go-scalability-talk/application/queues"
	"github.com/jadekler/git-go-scalability-talk/benchmark"
	"net"
	"os"
	"strings"
	"sync"
	"testing"
	"time"
)

var t tcpListenerBenchmark = tcpListenerBenchmark{}

type tcpListenerBenchmark struct {
	l    *listeners.TcpListener
	wg   *sync.WaitGroup
	q    queues.Queue
	p    int
	conn net.Conn
}

func TestMain(m *testing.M) {
	fmt.Println("Setup!")
	t.p = benchmark.GetOpenTcpPort()

	t.wg = &sync.WaitGroup{}
	t.q = benchmark.NewWaitingQueue(t.wg)

	t.l = listeners.NewTcpListener(t.p)
	go t.l.StartAccepting(t.q)
	t.conn = openTcpConn(t.p)

	os.Exit(m.Run())
}

func BenchmarkTcpListener(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			t.wg.Add(1)
			streamTcpItem(t.conn)
		}
	})

	t.wg.Wait()
}

func BenchmarkTcpListenerParallel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		t.wg.Add(1)
		streamTcpItem(t.conn)
	}

	t.wg.Wait()
}

// Open with some minimal retry
func openTcpConn(port int) net.Conn {
	var err error

	for i := 0; i < 5; i++ {
		conn, err := net.Dial("tcp", fmt.Sprintf("127.0.0.1:%d", port))

		if err != nil {
			if strings.Contains(err.Error(), "getsockopt: connection refused") {
				fmt.Println("Retrying conn")
				time.Sleep(100 * time.Millisecond)
				continue
			}

			panic(err)
		}

		return conn
	}

	panic(err)
}

func streamTcpItem(c net.Conn) {
	_, err := c.Write([]byte(benchmark.SMALL_MESSAGE))
	if err != nil {
		panic(err)
	}
}
