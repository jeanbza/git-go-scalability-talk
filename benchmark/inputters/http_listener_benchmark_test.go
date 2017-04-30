package benchmark

import (
	"bytes"
	"fmt"
	"github.com/jadekler/git-go-scalability-talk/application/inputters"
	"github.com/jadekler/git-go-scalability-talk/application/queues"
	"github.com/jadekler/git-go-scalability-talk/benchmark"
	"net/http"
	"sync"
	"testing"
)

var h httpListenerBenchmark = httpListenerBenchmark{}

type httpListenerBenchmark struct {
	l  *listeners.HttpListener
	wg *sync.WaitGroup
	q  queues.Queue
	p  int
}

func BenchmarkHttpListener(b *testing.B) {
	if h.l == nil {
		h.p = benchmark.GetOpenTcpPort()

		h.wg = &sync.WaitGroup{}
		h.q = benchmark.NewWaitingQueue(h.wg)

		h.l = listeners.NewHttpListener(h.p)
		go h.l.StartAccepting(h.q)
	}

	for i := 0; i < b.N; i++ {
		h.wg.Add(1)
		post(h.p)
	}

	h.wg.Wait()
}

func post(port int) {
	body := bytes.NewBufferString(LARGE_MESSAGE)
	http.Post(fmt.Sprintf("http://localhost:%d", port), "application/json", body)
}
