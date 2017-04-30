package queues

import "sync"

type MutexQueue struct {
	data [][]byte
	mu   *sync.Mutex
}

func NewMutexQueue() *MutexQueue {
	return &MutexQueue{
		mu: &sync.Mutex{},
	}
}

func (q *MutexQueue) Enqueue(data []byte) {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.data = append(q.data, data)
}

func (q *MutexQueue) Dequeue() ([]byte, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if len(q.data) == 0 {
		return nil, false
	}

	data := q.data[0]
	q.data = q.data[1:len(q.data)]
	return data, true
}
