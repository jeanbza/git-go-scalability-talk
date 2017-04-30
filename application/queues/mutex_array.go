package queues

import "sync"

type MutexArrayQueue struct {
	data [][]byte
	mu   *sync.Mutex
}

func NewMutexArrayQueue() *MutexArrayQueue {
	return &MutexArrayQueue{
		mu: &sync.Mutex{},
	}
}

func (q *MutexArrayQueue) Enqueue(data []byte) {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.data = append(q.data, data)
}

func (q *MutexArrayQueue) Dequeue() ([]byte, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if len(q.data) == 0 {
		return nil, false
	}

	data := q.data[0]
	q.data = q.data[1:len(q.data)]
	return data, true
}
