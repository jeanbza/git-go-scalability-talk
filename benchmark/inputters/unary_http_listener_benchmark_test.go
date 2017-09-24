package benchmark

import (
	"bytes"
	"fmt"
	"github.com/jadekler/git-go-scalability-talk/benchmark"
	"net/http"
	"testing"
)

func BenchmarkUnaryHttpListener(b *testing.B) {
	for i := 0; i < b.N; i++ {
		h.wg.Add(1)
		post(h.p)
	}

	h.wg.Wait()
}

func BenchmarkUnaryHttpListenerParallel(b *testing.B) {
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            h.wg.Add(1)
            post(h.p)
        }
    })

	h.wg.Wait()
}

func post(port int) {
	body := bytes.NewBufferString(benchmark.SMALL_MESSAGE)
	http.Post(fmt.Sprintf("http://localhost:%d", port), "application/json", body)
}
