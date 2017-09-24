package benchmark

import (
	"fmt"
	"github.com/jadekler/git-go-scalability-talk/benchmark"
	"net"
	"strings"
	"testing"
	"time"
)

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
