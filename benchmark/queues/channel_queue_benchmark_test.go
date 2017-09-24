package queues

import (
	"github.com/jadekler/git-go-scalability-talk/application/queues"
	"github.com/jadekler/git-go-scalability-talk/benchmark"
	"sync"
	"testing"
)

var c channelQueueBenchmark = channelQueueBenchmark{}

type channelQueueBenchmark struct {
	q  queues.Queue
	wg *sync.WaitGroup
}

func BenchmarkChannelQueue(b *testing.B) {
	if c.q == nil {
		c.q = queues.NewChannelQueue()
		c.wg = &sync.WaitGroup{}
		go constantlyDequeue(c.wg, c.q)
	}

	for i := 0; i < b.N; i++ {
		c.wg.Add(1)
		c.q.Enqueue([]byte(benchmark.VERY_LARGE_MESSAGE))
	}

	c.wg.Wait()
}
