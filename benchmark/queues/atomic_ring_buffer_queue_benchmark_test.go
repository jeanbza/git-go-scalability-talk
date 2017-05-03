// Based on: github.com/cloudfoundry/diodes/one_to_one.go
// Note that this solution has some issues as it approaches math.MaxInt64 related to the size
// modding, but it's probably not enough to care about
package queues

import (
	"github.com/jadekler/git-go-scalability-talk/application/queues"
	"github.com/jadekler/git-go-scalability-talk/benchmark"
	"sync"
	"testing"
)

var ar atomicRingBufferQueueBenchmark = atomicRingBufferQueueBenchmark{}

type atomicRingBufferQueueBenchmark struct {
	q  queues.Queue
	wg *sync.WaitGroup
}

func BenchmarkAtomicRingBufferQueue(b *testing.B) {
	if ar.q == nil {
		ar.q = queues.NewAtomicRingBuffer(10000)
		ar.wg = &sync.WaitGroup{}
		go constantlyDequeue(ar.wg, ar.q)
	}

	for i := 0; i < b.N; i++ {
		ar.wg.Add(1)
		ar.q.Enqueue([]byte(benchmark.VERY_LARGE_MESSAGE))
	}

	ar.wg.Wait()
}
