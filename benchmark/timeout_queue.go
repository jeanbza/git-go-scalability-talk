package benchmark

import (
	"sync"
	"time"
)

type TimeoutQueue struct {
	wg *sync.WaitGroup
	t  *time.Timer
	d  *time.Duration
}

func NewTimeoutQueue(wg *sync.WaitGroup, t *time.Timer, d *time.Duration) *TimeoutQueue {
	return &TimeoutQueue{
		wg: wg,
		t:  t,
		d:  d,
	}
}

func (q *TimeoutQueue) Enqueue(data []byte) {
	q.t.Reset(*q.d)
}

func (q *TimeoutQueue) Dequeue() ([]byte, bool) {
	return nil, false
}
