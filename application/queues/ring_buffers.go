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
    }
}

func (q *RingBuffer) Enqueue(data []byte) {
    q.mu.Lock()
    q.buffer[q.inputCursor] = data

    if (q.inputCursor == len(q.buffer) - 1) {
        q.inputCursor = -1
    }
    q.inputCursor++
    q.mu.Unlock()
}

func (q *RingBuffer) Dequeue() []byte {
    q.mu.Lock()
    data := q.buffer[q.outputCursor]

    if (q.outputCursor == len(q.buffer) - 1) {
        q.outputCursor = -1
    }
    q.outputCursor++
    q.mu.Unlock()

    return data
}
