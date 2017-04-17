package queues

type Channeler struct {
    c chan []byte
}

func NewChanneler() *Channeler {
    return &Channeler{
        c: make(chan []byte, 10),
    }
}

func (q *Channeler) Enqueue(data []byte) {
    q.c <- data
}

func (q *Channeler) Dequeue() ([]byte, bool) {
    select {
    case data := <-q.c:
        return data, true
    default:
        return nil, false
    }
}
