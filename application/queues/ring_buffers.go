package queues

import "sync"

type RingBuffer struct {
    buffer       [][]byte
    inputCursor  int
    outputCursor int
    mu           *sync.Mutex
}

func NewRingBuffer(size int) *RingBuffer {
    return &RingBuffer{
        buffer:       make([][]byte, size),
        inputCursor:  0,
        outputCursor: 0,
        mu:           &sync.Mutex{},
    }
}

func (q *RingBuffer) Enqueue(data []byte) {
    q.mu.Lock()
    defer q.mu.Unlock()

    q.buffer[q.inputCursor] = data

    if (q.inputCursor == len(q.buffer)-1) {
        q.inputCursor = -1
    }
    q.inputCursor++
}

func (q *RingBuffer) Dequeue() ([]byte, bool) {
    q.mu.Lock()
    defer q.mu.Unlock()

    if q.outputCursor == q.inputCursor {
        return nil, false
    }

    data := q.buffer[q.outputCursor]

    if (q.outputCursor == len(q.buffer)-1) {
        q.outputCursor = -1
    }
    q.outputCursor++

    return data, true
}
