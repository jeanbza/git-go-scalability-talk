package benchmark

import (
	"sync"
)

type WaitingQueue struct {
	wg *sync.WaitGroup
}

func NewWaitingQueue(wg *sync.WaitGroup) *WaitingQueue {
	return &WaitingQueue{
		wg: wg,
	}
}

func (q *WaitingQueue) Enqueue(data []byte) {
	q.wg.Done()
}

func (q *WaitingQueue) Dequeue() ([]byte, bool) {
	return nil, false
}
