package queues

import (
	"github.com/jadekler/git-go-scalability-talk/application/queues"
	"sync"
)

func constantlyDequeue(wg *sync.WaitGroup, q queues.Queue) {
	for {
		if _, ok := q.Dequeue(); ok {
			wg.Done()
		}
	}
}
