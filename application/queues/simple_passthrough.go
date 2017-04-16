package queues

import "sync"

type SimplePasser struct {
    data []byte
    mu   *sync.Mutex
}

func (q *SimplePasser) Enqueue(data []byte) {
    q.mu.Lock()
    q.data = data
}

func (q *SimplePasser) Dequeue() []byte {
    data := q.data
    q.mu.Unlock()
    return data
}
