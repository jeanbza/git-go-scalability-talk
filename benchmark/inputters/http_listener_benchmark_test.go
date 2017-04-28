package benchmark

import (
	"fmt"
	"github.com/jadekler/git-go-scalability-talk/application/inputters"
	"github.com/jadekler/git-go-scalability-talk/application/queues"
	"sync"
	"testing"
    "github.com/jadekler/git-go-scalability-talk/benchmark"
    "bytes"
    "net/http"
)

var (
	l  *listeners.HttpListener = nil
	wg *sync.WaitGroup         = nil
	q  queues.Queue            = nil
	p  int
)

func BenchmarkHttpListener(b *testing.B) {
	if l == nil {
		p = benchmark.GetOpenTcpPort()
		fmt.Println("Starting on port", p)

		wg = &sync.WaitGroup{}
		q = benchmark.NewWaitingQueue(wg)

		l = listeners.NewHttpListener(p)
		go l.StartAccepting(q)
	}

	for i := 0; i < b.N; i++ {
		wg.Add(1)
		post(p)
	}

	wg.Wait()
}

func post(port int) {
    body := bytes.NewBufferString("Hello world")
    http.Post(fmt.Sprintf("http://localhost:%d", port), "application/json", body)
}