package queues

import (
    "github.com/jadekler/git-go-scalability-talk/application/queues"
    "github.com/jadekler/git-go-scalability-talk/benchmark"
    "sync"
    "testing"
)

var ma mutexArrayQueueBenchmark = mutexArrayQueueBenchmark{}

type mutexArrayQueueBenchmark struct {
    q  queues.Queue
    wg *sync.WaitGroup
}

func BenchmarkMutexArrayQueue(b *testing.B) {
    if ma.q == nil {
        ma.q = queues.NewMutexArrayQueue()
        ma.wg = &sync.WaitGroup{}
        go constantlyDequeue(ma.wg, ma.q)
    }

    for i := 0; i < b.N; i++ {
        ma.wg.Add(1)
        ma.q.Enqueue([]byte(benchmark.VERY_LARGE_MESSAGE))
    }

    ma.wg.Wait()
}
