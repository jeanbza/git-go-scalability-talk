package queues

type Queue interface {
    Enqueue(data []byte)
    Dequeue() ([]byte, bool)
}
