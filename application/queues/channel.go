package queues

type ChannelQueue struct {
	c chan []byte
}

func NewChannelQueue() *ChannelQueue {
	return &ChannelQueue{
		c: make(chan []byte, 1000),
	}
}

func (q *ChannelQueue) Enqueue(data []byte) {
	select {
	case q.c <- data:
	default:
		// Queue was full! Data dropped
		// metrics.Record("dropped_messages", 1)
	}
}

func (q *ChannelQueue) Dequeue() ([]byte, bool) {
	select {
	case data := <-q.c:
		return data, true
	default:
		return nil, false
	}
}


