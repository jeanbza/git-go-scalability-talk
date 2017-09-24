package benchmark

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"testing"
)

func BenchmarkUnaryHttpListener(b *testing.B) {
	for i := 0; i < b.N; i++ {
		post(h.p, h.wg)
	}

	h.wg.Wait()
}

func BenchmarkUnaryHttpListenerParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			post(h.p, h.wg)
		}
	})

	h.wg.Wait()
}

func post(port int, wg *sync.WaitGroup) {
	wg.Add(1)

	body := bytes.NewBufferString(msg)
	_, err := http.Post(fmt.Sprintf("http://localhost:%d", port), "application/json", body)
	if err != nil {
		if strings.Contains(err.Error(), "read: connection reset by peer") {
			// server is closing connection when we try to talk to it; bummer about doing this in parallel at high speeds
			wg.Done() // that message never got sent, so decrement wg
			return
		}
		panic(err)
	}
}
