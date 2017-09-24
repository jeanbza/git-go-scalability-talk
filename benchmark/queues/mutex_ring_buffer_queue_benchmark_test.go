package queues

import (
	"github.com/jadekler/git-go-scalability-talk/application/queues"
	"github.com/jadekler/git-go-scalability-talk/benchmark"
	"sync"
	"testing"
)

var mr mutexRingBufferQueueBenchmark = mutexRingBufferQueueBenchmark{}

type mutexRingBufferQueueBenchmark struct {
	q  queues.Queue
	wg *sync.WaitGroup
}

func BenchmarkMutexRingBufferQueue(b *testing.B) {
	if mr.q == nil {
		mr.q = queues.NewMutexRingBufferQueue(10000)
		mr.wg = &sync.WaitGroup{}
		go constantlyDequeue(mr.wg, mr.q)
	}

	for i := 0; i < b.N; i++ {
		mr.wg.Add(1)
		mr.q.Enqueue([]byte(benchmark.VERY_LARGE_MESSAGE))
	}

	mr.wg.Wait()
}
